package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/transaction"
)

//contractDeploymentTimeout  is the default timeout for contract deployment
const contractDeploymentTimeout = 20 * time.Second
const contractCreationGasLimit uint64 = 1000000
const contractCallMethodGasLimit uint64 = 500000
const transferGasLimit uint64 = 1000000

//ErrorUnknownContractType indicates an unsupported contract type
var ErrorUnknownContractType = errors.New("unknown contract type")

//ErrorMethodNameNotFound indicates the method to call of a contract does not exist
var ErrorMethodNameNotFound = errors.New("contract method name not found")

// ChainEthereum is a wrapper for ethereum client
type ChainEthereum struct {
	contractBackend    bind.ContractBackend
	deployBackend      bind.DeployBackend
	contractEvents     chan *chain.ContractEvent
	chainStateReader   ethereum.ChainStateReader
	pendingStateReader ethereum.PendingStateReader
}

// NewChainEthereum creates a new Ethereum chain object with an existing ethclient
func NewChainEthereum(client *ethclient.Client) *ChainEthereum {
	return &ChainEthereum{
		contractBackend:    client,
		deployBackend:      client,
		contractEvents:     make(chan *chain.ContractEvent, 1024),
		chainStateReader:   client,
		pendingStateReader: client,
	}
}

// NewSimChainEthereum creates a new Ethereum chain object with the SimulatedBackend
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
func (c *ChainEthereum) NewContract(ctx context.Context, from []byte, contractType string, contractData interface{}) (*transaction.ContractTransaction, error) {
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

// Call inits a transaction to call a contract method
// The transaction is not sent out yet and must be confirmed later by sender
// param "value" is the money to sent to the transaction address
// param "callData" is a container of all the args needed for method
func (c *ChainEthereum) Call(ctx context.Context, from []byte, contractType string, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.ContractTransaction, error) {
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

func (c *ChainEthereum) callMethod(ctx context.Context, from []byte, abiParsed abi.ABI, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.ContractTransaction, error) {
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

	ct := transaction.NewContractTransaction(rawTx, from)
	ct.ContractAddr = contractAddr
	ct.TxHashFunc = func(rawTx interface{}) []byte {
		return types.HomesteadSigner{}.Hash(rawTx.(*types.Transaction)).Bytes()
	}
	ct.TxHexHashSignedFunc = func(rawTx interface{}) string {
		return rawTx.(*types.Transaction).Hash().Hex()
	}
	ct.SignFunc = func(sig []byte) error {
		cpyTx, err := ct.RawTx().(*types.Transaction).WithSignature(types.HomesteadSigner{}, sig)
		if err != nil {
			return fmt.Errorf("failed to update transaction signature:%v", err)
		}

		ct.SetRawTx(cpyTx)
		return nil
	}
	return ct, nil
}

func (c *ChainEthereum) createContract(ctx context.Context, from []byte, abi abi.ABI, bin string, contractData interface{}) (*transaction.ContractTransaction, error) {
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
	// gasLimit, err = c.contractBackend.EstimateGas(ctx, msg)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	// }

	//use a hardcoded gasLimit temporarily
	rawTx := types.NewContractCreation(nonce, value, contractCreationGasLimit /*asLimit*/, gasPrice, bytecode)

	ct := transaction.NewContractTransaction(rawTx, from)
	ct.ContractAddr = ethcrypto.CreateAddress(fromAddr, rawTx.Nonce()).Bytes()
	ct.TxHashFunc = func(rawTx interface{}) []byte {
		return types.HomesteadSigner{}.Hash(rawTx.(*types.Transaction)).Bytes()
	}
	ct.TxHexHashSignedFunc = func(rawTx interface{}) string {
		return rawTx.(*types.Transaction).Hash().Hex()
	}
	ct.SignFunc = func(sig []byte) error {
		cpyTx, err := ct.RawTx().(*types.Transaction).WithSignature(types.HomesteadSigner{}, sig)
		if err != nil {
			return fmt.Errorf("failed to update transaction signature:%v", err)
		}

		ct.SetRawTx(cpyTx)

		return nil
	}

	return ct, nil
}

//ConfirmTrans update the raw transaction with given signature and send it to the underlying block chain.
func (c *ChainEthereum) ConfirmTrans(ctx context.Context, trans *transaction.ContractTransaction, sig []byte) error {

	if err := trans.WithSign(sig); err != nil {
		return err
	}
	rawTx := trans.RawTx().(*types.Transaction)
	sender, err := types.HomesteadSigner{}.Sender(rawTx)
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
func (c *ChainEthereum) WaitMined(ctx context.Context, trans *transaction.ContractTransaction) error {
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
					Addr: addr,
					Name: eventName,
					T:    eventVType,
					V:    v,
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

func (c *ChainEthereum) BalanceAt(ctx context.Context, addr []byte) (*big.Int, error) {

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

// UnmarshalPubkey converts base64 string to a secp256k1 public key.
func (c *ChainEthereum) UnmarshalPubkey(pub string) (*ecdsa.PublicKey, error) {

	buf, err := base64.StdEncoding.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	pubKey := ethcrypto.ToECDSAPub(buf)
	return pubKey, nil
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
func (c *ChainEthereum) Transfer(ctx context.Context, from []byte, to []byte, value *big.Int, prv *ecdsa.PrivateKey) (err error) {
	fromAddr := common.BytesToAddress(from)
	toAddr := common.BytesToAddress(to)
	fmt.Printf("fromAddr, to, value: %v %v %v\n", fromAddr.Hex(), toAddr.Hex(), value)
	var nonce uint64
	if nonce, err = c.pendingStateReader.PendingNonceAt(ctx, fromAddr); err != nil {
		return fmt.Errorf("failed to retrieve account nonce: %v", err)
	}
	var gasPrice *big.Int
	if gasPrice, err = c.contractBackend.SuggestGasPrice(ctx); err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	// new transaction
	rawTx := types.NewTransaction(nonce, toAddr, value, transferGasLimit, gasPrice, nil)
	fmt.Printf("Cost %v\n", rawTx.Cost())

	// sign
	signer := types.HomesteadSigner{}
	h := signer.Hash(rawTx)
	sig, err := c.Sign(h[:], prv)
	if err != nil {
		return err
	}

	// sign rawTransaction with sig
	signedTx, err := rawTx.WithSignature(signer, sig)
	// sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return err
	}

	// submit to txpool
	err = c.contractBackend.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	// wait to mine
	receipt, err := bind.WaitMined(ctx, c.deployBackend, signedTx)
	if err != nil {
		return fmt.Errorf("wait mined returns error:%v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Printf("transaction receipt address:%v, status:%v\n", receipt.ContractAddress.Hex(), receipt.Status)
	}
	return err
}

//Transfer token to other address
func (c *ChainEthereum) TransferICST(ctx context.Context, from []byte, to []byte, value *big.Int) (*transaction.ContractTransaction, error) {
	callData := map[string]interface{}{
		"_to":    common.BytesToAddress(to),
		"_value": value,
	}
	return c.Call(ctx, from, "ICST", common.Hex2Bytes(ICSTAddr), "transfer", big.NewInt(0), callData)
}
