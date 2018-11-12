package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/chain"
	icstcommon "github.com/icstglobal/go-icst/common"
	"github.com/icstglobal/go-icst/transaction"
)

//contractDeploymentTimeout  is the default timeout for contract deployment
const contractDeploymentTimeout = 20 * time.Second
const contractCreationGasLimit uint64 = 1000000
const contractCallMethodGasLimit uint64 = 500000
const transferGasLimit uint64 = 1000000

// abi cache
var abiCache = map[[32]byte]abi.ABI{}

//ErrorUnknownContractType indicates an unsupported contract type
var ErrorUnknownContractType = errors.New("unknown contract type")

//ErrorMethodNameNotFound indicates the method to call of a contract does not exist
var ErrorMethodNameNotFound = errors.New("contract method name not found")

// ChainEthereum is a wrapper for ethereum client
type ChainEthereum struct {
	contractBackend  bind.ContractBackend
	deployBackend    bind.DeployBackend
	contractEvents   chan *chain.ContractEvent
	chainStateReader ethereum.ChainStateReader
	chainReader      ethereum.ChainReader
}

// NewChainEthereum creates a new Ethereum chain object with an existing ethclient
func NewChainEthereum(client *ethclient.Client) *ChainEthereum {
	return &ChainEthereum{
		contractBackend:  client,
		deployBackend:    client,
		chainStateReader: client,
		chainReader:      client,
		contractEvents:   make(chan *chain.ContractEvent, 1024),
	}
}

//NewSimChainEthereum creates a new Ethereum chain object with the SimulatedBackend
//Don't support ChainReader interface
func NewSimChainEthereum(backend *backends.SimulatedBackend) *ChainEthereum {
	return &ChainEthereum{
		contractBackend:  backend,
		deployBackend:    backend,
		contractEvents:   make(chan *chain.ContractEvent, 1024),
		chainStateReader: backend,
	}
}

//DialEthereum creates a new Ethereum chain object by dialing to the given url
func DialEthereum(url string) (*ChainEthereum, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return NewChainEthereum(client), nil
}

// GetContract gets smart contract from Ethereum chain with its address.
func (c *ChainEthereum) GetContract(addr []byte, contractType string) (interface{}, error) {
	switch contractType {
	case "Content":
		return c.getContentContract(addr)
	case "Skill":
		return c.getSkillContract(addr)
	case "ICST":
		return c.getICSTContract(addr)
	default:
		return nil, fmt.Errorf("unknown contract type:%v", contractType)
	}
}

//NewContract makes a new contract creation transaction.
//The transaction is not sent out yet and must be confirmed later by sender.
func (c *ChainEthereum) NewContract(ctx context.Context, from []byte, contractType string, contractData interface{}) (*transaction.Transaction, error) {
	var parsed abi.ABI
	var bin string
	var err error
	switch contractType {
	case "Content":
		parsed, err = abi.JSON(strings.NewReader(ConsumeContentABI))
		bin = ConsumeContentBin
		break
	case "Skill":
		parsed, err = abi.JSON(strings.NewReader(ConsumeSkillABI))
		bin = ConsumeSkillBin
		break
	case "ICST":
		parsed, err = abi.JSON(strings.NewReader(ICSTABI))
		bin = ICSTBin
		break
	default:
		err = fmt.Errorf("unknown contract type:%v", contractType)
	}

	if err != nil {
		return nil, err
	}

	return c.createContract(ctx, from, parsed, bin, contractData)
}

// Call without send transaction
func (c *ChainEthereum) Query(ctx context.Context, addr []byte, abiString string, methodName string, result interface{}, params ...interface{}) error {
	abiParsed, err := getAbiFromCache(abiString)
	if err != nil {
		return err
	}

	ctr := bind.NewBoundContract(common.BytesToAddress(addr), abiParsed, c.contractBackend, c.contractBackend, c.contractBackend)
	return ctr.Call(nil, result, methodName, params...)
}

// Call inits a transaction to call a contract method
// The transaction is not sent out yet and must be confirmed later by sender
// param "value" is the money to sent to the transaction address
// param "callData" is a container of all the args needed for method
func (c *ChainEthereum) Call(ctx context.Context, from []byte, contractType string, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.Transaction, error) {
	var abiParsed abi.ABI
	var err error
	switch contractType {
	case "Content":
		if abiParsed, err = abi.JSON(strings.NewReader(ConsumeContentABI)); err != nil {
			return nil, err
		}
		break
	case "Skill":
		if abiParsed, err = abi.JSON(strings.NewReader(ConsumeSkillABI)); err != nil {
			return nil, err
		}
		break
	case "ICST":
		if abiParsed, err = abi.JSON(strings.NewReader(ICSTABI)); err != nil {
			return nil, err
		}
		break
	default:
		return nil, ErrorUnknownContractType
	}
	return c.callMethod(ctx, from, abiParsed, contractAddr, methodName, value, callData)
}

// CallWithAbi inits a transaction to call a contract method
// The transaction is not sent out yet and must be confirmed later by sender
// param "value" is the money to sent to the transaction address
// param "callData" is a container of all the args needed for method
func (c *ChainEthereum) CallWithAbi(ctx context.Context, from []byte, contractAddr []byte, methodName string, value *big.Int, callData interface{}, abiStr string) (*transaction.Transaction, error) {
	_abi, err := getAbiFromCache(abiStr)
	if err != nil {
		return nil, err
	}
	return c.callMethod(ctx, from, _abi, contractAddr, methodName, value, callData)
}

func (c *ChainEthereum) callMethod(ctx context.Context, from []byte, abiParsed abi.ABI, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.Transaction, error) {
	method, found := abiParsed.Methods[methodName]
	if !found {
		return nil, ErrorMethodNameNotFound
	}
	params, err := extractAbiParams(method, callData)
	if err != nil {
		return nil, err
	}
	input, err := abiParsed.Pack(methodName, params...)
	if err != nil {
		return nil, err
	}
	fromAddr := common.BytesToAddress(from)
	var nonce uint64
	if nonce, err = c.contractBackend.PendingNonceAt(ctx, fromAddr); err != nil {
		return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
	}
	var gasPrice *big.Int
	if gasPrice, err = c.contractBackend.SuggestGasPrice(ctx); err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	rawTx := types.NewTransaction(nonce, common.BytesToAddress(contractAddr), value, contractCallMethodGasLimit, gasPrice, input)

	ct := transaction.NewTransaction(rawTx, from)
	ct.To = contractAddr
	ct.TxHashFunc = func(rawTx interface{}) []byte {
		return signer().Hash(rawTx.(*types.Transaction)).Bytes()
	}
	ct.TxHexHashSignedFunc = func(rawTx interface{}) string {
		return rawTx.(*types.Transaction).Hash().Hex()
	}
	ct.SignFunc = func(sig []byte) error {
		cpyTx, err := ct.RawTx().(*types.Transaction).WithSignature(signer(), sig)
		if err != nil {
			return fmt.Errorf("failed to update transaction signature:%v", err)
		}

		ct.SetRawTx(cpyTx)
		return nil
	}
	return ct, nil
}

func (c *ChainEthereum) createContract(ctx context.Context, from []byte, abi abi.ABI, bin string, contractData interface{}) (*transaction.Transaction, error) {
	params, err := extractAbiParams(abi.Constructor, contractData)
	if err != nil {
		return nil, err
	}
	// empty method name means "constructor" method
	input, err := abi.Pack("", params...)
	if err != nil {
		return nil, err
	}
	bytecode := common.FromHex(bin)
	bytecode = append(bytecode, input...)
	fromAddr := common.BytesToAddress(from)
	var nonce uint64
	if nonce, err = c.contractBackend.PendingNonceAt(ctx, fromAddr); err != nil {
		return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
	}
	value := new(big.Int)
	var gasPrice *big.Int
	if gasPrice, err = c.contractBackend.SuggestGasPrice(ctx); err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	//TODO: gasLimit estimated does not enough for contract creation.
	// var gasLimit uint64
	// msg := ethereum.CallMsg{From: fromAddr, To: nil, Value: value, Data: input}
	// gasLimit, err = c.clientInterface.EstimateGas(ctx, msg)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	// }

	//use a hardcoded gasLimit temporarily
	rawTx := types.NewContractCreation(nonce, value, contractCreationGasLimit /*asLimit*/, gasPrice, bytecode)

	ct := transaction.NewTransaction(rawTx, from)
	ct.To = ethcrypto.CreateAddress(fromAddr, rawTx.Nonce()).Bytes()
	ct.TxHashFunc = func(rawTx interface{}) []byte {
		return signer().Hash(rawTx.(*types.Transaction)).Bytes()
	}
	ct.TxHexHashSignedFunc = func(rawTx interface{}) string {
		return rawTx.(*types.Transaction).Hash().Hex()
	}
	ct.SignFunc = func(sig []byte) error {
		cpyTx, err := ct.RawTx().(*types.Transaction).WithSignature(signer(), sig)
		if err != nil {
			return fmt.Errorf("failed to update transaction signature:%v", err)
		}

		ct.SetRawTx(cpyTx)

		return nil
	}

	return ct, nil
}

//ConfirmTrans update the raw transaction with given signature and send it to the underlying block chain.
func (c *ChainEthereum) ConfirmTrans(ctx context.Context, trans *transaction.Transaction, sig []byte) error {

	if err := trans.WithSign(sig); err != nil {
		return err
	}
	rawTx := trans.RawTx().(*types.Transaction)
	sender, err := signer().Sender(rawTx)
	if err != nil {
		return fmt.Errorf("failed to recover sender's address from the signature: %v", err)
	}
	if !reflect.DeepEqual(sender.Bytes(), trans.Sender()) {
		return fmt.Errorf("sender's address does not match")
	}
	return c.contractBackend.SendTransaction(ctx, rawTx)
}

//extractAbiParams retrieves the values for arguments of abi's method, from the contractData object.
//names of arguments and fields of the contractData will be matched ignoring case.
func extractAbiParams(method abi.Method, contractData interface{}) ([]interface{}, error) {
	params := make([]interface{}, 0, len(method.Inputs))
	obj := contractData.(map[string]interface{})

	for _, arg := range method.Inputs {
		params = append(params, obj[arg.Name])
	}
	return params, nil
}

// getContentContract gets content contract from Ethereum chain with its address.
func (c *ChainEthereum) getContentContract(addr []byte) (*ConsumeContent, error) {
	ct, err := NewConsumeContent(common.BytesToAddress(addr), c.contractBackend)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// getSkillContract gets skill contract from Ethereum chain with its address.
func (c *ChainEthereum) getSkillContract(addr []byte) (*ConsumeSkill, error) {
	ct, err := NewConsumeSkill(common.BytesToAddress(addr), c.contractBackend)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

// getICSTContract gets skill contract from Ethereum chain with its address.
func (c *ChainEthereum) getICSTContract(addr []byte) (*ICST, error) {
	ct, err := NewICST(common.BytesToAddress(addr), c.contractBackend)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

//WaitMined blocks the caller until the transaction is mined, or gets an error
func (c *ChainEthereum) WaitMined(ctx context.Context, trans *transaction.Transaction) error {
	receipt, err := bind.WaitMined(ctx, c.deployBackend, trans.RawTx().(*types.Transaction))
	if err != nil {
		return fmt.Errorf("wait mined returns error:%v", err)
	}

	//TODO: check what exactly "receipt.Status" means.
	//When creating contract, we always get a "0" as the staus, but the contract can be sucessly deployed and called.
	if receipt.Status == types.ReceiptStatusFailed {
		// return fmt.Errorf("transaction failed")
		log.Printf("transaction receipt address:%v, status:%v\n", receipt.ContractAddress.Hex(), receipt.Status)
	}
	return err
}

//WaitContractDeployed blocks the caller untile the transactio to create a contract is mined, or gets an error.
//The difference with WaitMined is that it also make sure the contrace code is not empty.
func (c *ChainEthereum) WaitContractDeployed(ctx context.Context, tx interface{}) (common.Address, error) {
	return bind.WaitDeployed(ctx, c.deployBackend, tx.(*types.Transaction))
}

func (c *ChainEthereum) watchEvent(ctx context.Context, contractDeployed *ConsumeSkill, stateChan chan<- *ConsumeSkillStateChange) (event.Subscription, error) {
	watchOpts := &bind.WatchOpts{Start: nil, Context: ctx} // start from the latest block
	return contractDeployed.WatchStateChange(watchOpts, stateChan)
}

//WatchContractEvent listening on the events from contract, and wrap it in a general contract event struct
//It returns error if the given event if not found by name
func (c *ChainEthereum) WatchContractEvent(ctx context.Context, addr []byte, contractType string, eventName string, eventVType reflect.Type) (<-chan *chain.ContractEvent, error) {
	abi, err := getAbi(contractType)
	if err != nil {
		return nil, err
	}

	ctr := bind.NewBoundContract(common.BytesToAddress(addr), abi, c.contractBackend, c.contractBackend, c.contractBackend)
	opts := new(bind.WatchOpts)
	opts.Context = ctx
	//watch from the latest block
	logs, sub, err := ctr.WatchLogs(opts, eventName)
	if err != nil {
		return nil, err
	}

	quit := make(chan struct{}, 1)
	//start event loop to convert eth logs to our contract event type
	go func() {
		for {
			select {
			case rawLog := <-logs:

				v := reflect.New(eventVType).Interface()
				if err = abi.Events[eventName].Inputs.Unpack(v, rawLog.Data); err != nil {
					// if err = unpack(abi.Events[eventName], v, rawLog.Data); err != nil {
					log.Println("[ERROR]failed to parse raw event log,", err)
					break
				}

				c.contractEvents <- &chain.ContractEvent{
					Addr:      addr,
					Name:      eventName,
					T:         eventVType,
					V:         v,
					BlockNum:  rawLog.BlockNumber,
					BlockHash: rawLog.BlockHash[:],
					TxIndex:   uint64(rawLog.TxIndex),
					TxHash:    rawLog.TxHash[:],
					Unwatch: func() {
						var q struct{}
						quit <- q         //quit the event loop
						sub.Unsubscribe() //unsubscribe event watching on block chain
					},
				}
			case <-ctx.Done():
				break
			case <-quit:
				break
			}
		}
	}()

	return c.contractEvents, nil
}

// GetContractEvents returns multiple contract events of the same contract together, with the same order inside block
// This method ensures that all events needed will be returned, or none of them will be returned if an error happen.
func (c *ChainEthereum) GetContractEvents(ctx context.Context, addr []byte, fromBlock, toBlock *big.Int, abiString string, eventTypes map[string]reflect.Type) ([]*chain.ContractEvent, error) {
	abiParsed, err := getAbiFromCache(abiString)
	if err != nil {
		return nil, err
	}

	eventTopicMap := make(map[string]common.Hash)
	topics := make([][]common.Hash, 0)
	topic := make([]common.Hash, 0)
	// query "all" matching events of the same contract, by using topic
	// {{EventA, EventB}}, instead of
	// {{EventA}, {EventB}}
	for eventName := range eventTypes {
		eventID := abiParsed.Events[eventName].Id()
		topic = append(topic, eventID)
		// track topic for later log parsing
		eventTopicMap[eventName] = eventID
	}
	topics = append(topics, topic)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.BytesToAddress(addr)},
		Topics:    topics,
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	}
	logs, err := c.contractBackend.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	events := make([]*chain.ContractEvent, 0)
	ctr := bind.NewBoundContract(common.BytesToAddress(addr), abiParsed, c.contractBackend, c.contractBackend, c.contractBackend)
	for _, rawLog := range logs {
		var evt *chain.ContractEvent
		// try event types one by one to see if the log can be pased
		for eventName, eventVType := range eventTypes {
			eventTopic := eventTopicMap[eventName]
			if !bytes.Equal(rawLog.Topics[0][:], eventTopic[:]) {
				//try next event type
				continue
			}
			v := reflect.New(eventVType).Interface()
			if err = ctr.UnpackLog(v, eventName, rawLog); err != nil {
				log.Println("[ERROR]failed to parse raw event log,", err)
				return nil, err
			}

			evt = &chain.ContractEvent{
				Addr:      addr,
				Name:      eventName,
				T:         eventVType,
				V:         v,
				BlockNum:  rawLog.BlockNumber,
				BlockHash: rawLog.BlockHash[:],
				TxIndex:   uint64(rawLog.TxIndex),
				TxHash:    rawLog.TxHash[:],
			}
		}
		// can not be parsed to any event type
		if evt == nil {
			log.Printf("[ERROR]unexpected event log topic,topic hex:%v, err:%v\n", rawLog.Topics[0].String(), err)
			return nil, err
		}

		events = append(events, evt)
	}

	return events, nil
}

//BalanceAt returns the balance of an account
func (c *ChainEthereum) BalanceAt(ctx context.Context, addr []byte) (*big.Int, error) {
	b, err := c.chainStateReader.BalanceAt(ctx, common.BytesToAddress(addr), nil)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//BalanceAtICST returns the ICST balance of an account
func (c *ChainEthereum) BalanceAtICST(ctx context.Context, addr []byte) (*big.Int, error) {
	cxToken, err := c.GetContract(common.Hex2Bytes(ICSTAddr), "ICST")
	if err != nil {
		return nil, err
	}
	b, err := cxToken.(*ICST).BalanceOf(nil, common.BytesToAddress(addr))

	if err != nil {
		return nil, err
	}
	return b, nil
}

//PubKeyToAddress convert public key to chain specific address
func (c *ChainEthereum) PubKeyToAddress(pub *ecdsa.PublicKey) []byte {
	return ethcrypto.PubkeyToAddress(*pub).Bytes()
}

func getAbi(contractType string) (abi.ABI, error) {
	var abiParsed abi.ABI
	var err error
	switch contractType {
	case "Content":
		if abiParsed, err = abi.JSON(strings.NewReader(ConsumeContentABI)); err != nil {
			return abi.ABI{}, err
		}
		break
	case "Skill":
		if abiParsed, err = abi.JSON(strings.NewReader(ConsumeSkillABI)); err != nil {
			return abi.ABI{}, err
		}
		break
	case "ICST":
		if abiParsed, err = abi.JSON(strings.NewReader(ICSTABI)); err != nil {
			return abi.ABI{}, err
		}
		break
	default:
		return abi.ABI{}, ErrorUnknownContractType
	}
	return abiParsed, nil
}

func getAbiFromStr(abiStr string) (abi.ABI, error) {
	var abiParsed abi.ABI
	var err error
	if abiParsed, err = abi.JSON(strings.NewReader(abiStr)); err != nil {
		return abi.ABI{}, err
	}
	return abiParsed, nil
}

// UnmarshalPubkey converts base64 string to a secp256k1 public key.
func (c *ChainEthereum) UnmarshalPubkey(pub string) (*ecdsa.PublicKey, error) {

	buf, err := base64.StdEncoding.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	return ethcrypto.UnmarshalPubkey(buf)
}

//MarshalPubKey convert a public key to base 64 string
func (c *ChainEthereum) MarshalPubKey(pub *ecdsa.PublicKey) string {
	buf := ethcrypto.FromECDSAPub(pub)
	return base64.StdEncoding.EncodeToString(buf)
}

//GenerateKey generates a new ecdsa public/private key pair
func (c *ChainEthereum) GenerateKey(ctx context.Context) (*ecdsa.PrivateKey, error) {
	return ethcrypto.GenerateKey()
}

//Sign data with privatekey
func (c *ChainEthereum) Sign(hash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	return ethcrypto.Sign(hash, prv)
}

//Transfer to other address
func (c *ChainEthereum) Transfer(ctx context.Context, from []byte, to []byte, value *big.Int) (*transaction.Transaction, error) {
	fromAddr := common.BytesToAddress(from)
	toAddr := common.BytesToAddress(to)
	fmt.Printf("fromAddr, to, value: %v %v %v\n", fromAddr.Hex(), toAddr.Hex(), value)
	nonce, err := c.contractBackend.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
	}
	gasPrice, err := c.contractBackend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	// new transaction
	rawTx := types.NewTransaction(nonce, toAddr, value, transferGasLimit, gasPrice, nil)
	ct := transaction.NewTransaction(rawTx, from)
	ct.To = to
	ct.TxHashFunc = func(rawTx interface{}) []byte {
		return signer().Hash(rawTx.(*types.Transaction)).Bytes()
	}
	ct.TxHexHashSignedFunc = func(rawTx interface{}) string {
		return rawTx.(*types.Transaction).Hash().Hex()
	}
	ct.SignFunc = func(sig []byte) error {
		cpyTx, err := ct.RawTx().(*types.Transaction).WithSignature(signer(), sig)
		if err != nil {
			return fmt.Errorf("failed to update transaction signature:%v", err)
		}

		ct.SetRawTx(cpyTx)
		return nil
	}
	fmt.Printf("Cost %v\n", rawTx.Cost())
	return ct, nil
}

//TransferICST token to other address
func (c *ChainEthereum) TransferICST(ctx context.Context, from []byte, to []byte, value *big.Int) (*transaction.Transaction, error) {
	callData := map[string]interface{}{
		"_to":    common.BytesToAddress(to),
		"_value": value,
	}
	return c.Call(ctx, from, "ICST", common.Hex2Bytes(ICSTAddr), "transfer", big.NewInt(0), callData)
}

//WatchBlocks loops for the blocks starting from the blockStart number. Order of blocks is guaranteed.
//Loops do not break even if there is some error happen
//Params:
// 	blockStart: block number to start from; nil means start from latest block
func (c *ChainEthereum) WatchBlocks(ctx context.Context, blockStart *big.Int) (<-chan *transaction.Block, <-chan error) {
	errors := make(chan error, 1)
	blocks := make(chan *transaction.Block, 8)
	hsginer := signer()

	go func() {
		incr := big.NewInt(1)
		start := blockStart
		for {
			//wait for 5s, as ethereum needs time to generate new blocks
			time.Sleep(time.Second * 5)

			select {
			case <-ctx.Done():
				break
			default:
				//get latest header
				latest, err := c.chainReader.HeaderByNumber(ctx, nil)
				if err != nil {
					errors <- err
					continue
				}
				// start from the latest block
				if start == nil {
					start = new(big.Int)
					*start = *latest.Number
				}
				for start.Cmp(latest.Number) <= 0 {
					rawBlock, err := c.chainReader.BlockByNumber(ctx, start)
					if err != nil {
						errors <- err
						break
					}
					block, err := parseBlockData(hsginer, rawBlock)
					if err != nil {
						errors <- err
						break
					}
					blocks <- block
					//increase only after the block is processed successfully
					start.Add(start, incr)
				}
			}
		}
	}()
	return blocks, errors
}

//WatchICSTTransfer loops for the blocks just like WatchBlocks, but filter out block's transactions not transferring ICST.
//Transactions in block are also different, `to` and `amount` has been change to the ICST receiver and amount.
func (c *ChainEthereum) WatchICSTTransfer(ctx context.Context, blockStart *big.Int) (<-chan *transaction.Block, <-chan error) {
	icstErrors := make(chan error, 1)
	icstBlocks := make(chan *transaction.Block, 8)
	icstAbi, _ := abi.JSON(strings.NewReader(ICSTABI))

	blocks, errors := c.WatchBlocks(ctx, blockStart)
	go func() {
		for {
			select {
			case b := <-blocks:
				nb := &transaction.Block{}
				nb.BlockNumber = b.BlockNumber
				nb.Hash = b.Hash
				nb.Trans = b.Trans[:0]
				for _, tx := range b.Trans {
					if tx.To == nil || len(tx.To) == 0 || len(tx.Data) == 0 || hex.EncodeToString(tx.To) != ICSTAddr {
						//ignore any transaction not related to ICST ERC20 token transfer
						continue
					}

					//check if call icst.Transfer
					method, err := icstAbi.MethodById(tx.Data)
					if err != nil || method.Name != "transfer" {
						//ignore
						continue
					}

					args, err := method.Inputs.UnpackValues(tx.Data[4:])
					if err != nil {
						icstErrors <- fmt.Errorf("failed to parse transfer args, tx.Hash:%v, inner error:%v", tx.Hash, err)
						continue
					}

					receipt, err := c.deployBackend.TransactionReceipt(ctx, common.HexToHash(tx.Hash))
					if err != nil {
						icstErrors <- fmt.Errorf("failed to get tx receipt, tx.Hash:%v, inner error:%v", tx.Hash, err)
						continue
					}

					to := args[0].(common.Address)
					amount := args[1].(*big.Int)
					ntx := &transaction.Message{}
					ntx.Hash = tx.Hash
					ntx.From = tx.From
					ntx.To = to.Bytes()
					ntx.Amount = amount
					ntx.CheckNonce = tx.CheckNonce
					ntx.GasLimit = tx.GasLimit
					ntx.GasPrice = tx.GasPrice
					ntx.Nonce = tx.Nonce
					ntx.Success = (receipt.Status == 1)
					ntx.GasUsed = receipt.GasUsed

					nb.Trans = append(nb.Trans, ntx)
				}

				icstBlocks <- nb
			case err := <-errors:
				icstErrors <- err
			}
		}
	}()
	return icstBlocks, icstErrors
}

func signer() types.Signer {
	return types.HomesteadSigner{}
}

func parseBlockData(s types.Signer, rawBlock *types.Block) (*transaction.Block, error) {
	block := &transaction.Block{}
	block.BlockNumber = rawBlock.NumberU64()
	block.Hash = rawBlock.Hash().Hex()

	for _, tx := range rawBlock.Transactions() {
		var msg types.Message
		var err error
		if tx.Protected() {
			msg, err = tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
		} else {
			msg, err = tx.AsMessage(s)
		}

		if err != nil {
			return nil, fmt.Errorf("transaction can not be parsed, block id:%v,tx hash:%v, inner err:%v", block.BlockNumber,
				tx.Hash().Hex(), err)
		}
		tm := &transaction.Message{}
		tm.Hash = tx.Hash().Hex()
		tm.Amount = msg.Value()
		tm.CheckNonce = msg.CheckNonce()
		tm.Data = msg.Data()
		tm.From = msg.From().Bytes()
		if msg.To() != nil {
			tm.To = msg.To().Bytes()
		}
		tm.GasLimit = msg.Gas()
		tm.GasPrice = msg.GasPrice()
		tm.Nonce = msg.Nonce()
		block.Trans = append(block.Trans, tm)
	}
	return block, nil
}

func getAbiFromCache(abiStr string) (abi.ABI, error) {
	abiHash := icstcommon.Hash([]byte(abiStr))
	var _abi abi.ABI
	_abi, ok := abiCache[abiHash]
	if ok {
		return _abi, nil
	}
	new_abi, err := getAbiFromStr(abiStr)
	if err != nil {
		return abi.ABI{}, err
	}
	abiCache[abiHash] = new_abi
	return new_abi, nil
}
