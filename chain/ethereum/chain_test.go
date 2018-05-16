package ethereum

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/icstglobal/go-icst/user"
)

func TestConsumeSkillContract(t *testing.T) {

	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey)
	t.Logf("ownerAddr:%v", ownerAddr.String())

	alloc := make(core.GenesisAlloc)
	for _, u := range []*user.User{owner} {
		alloc[crypto.PubkeyToAddress(u.PrivateKey.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	}

	simBackend := backends.NewSimulatedBackend(alloc)
	t.Log("sim backend created")
	chain := &ChainEthereum{contractBackend: simBackend, deployBackend: simBackend}
	contractData := struct {
		PHash      string
		PPublisher []byte //common.Address
		PPlatform  []byte //common.Address
		PConsumer  []byte //common.Address
		PPrice     uint32
		PRatio     uint8
	}{
		PHash:      "hash",
		PPublisher: ownerAddr.Bytes(),
		PPlatform:  ownerAddr.Bytes(),
		PConsumer:  ownerAddr.Bytes(),
		PPrice:     1,
		PRatio:     50,
	}
	//start a miner goroutine
	go func() {
		//mine after 2 seconds
		time.Sleep(2 * time.Second)
		simBackend.Commit()
		t.Log("contract deployment mined")
	}()
	t.Log("deploy contract")
	ct, err := chain.NewContract(context.Background(), ownerAddr.Bytes(), "Skill", contractData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("contract deployment transaction created")

	//sign the transaction locally, without send private key to the remote
	sig, err := crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction signed by sender")
	err = chain.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction sent to block chain")
	err = chain.WaitMined(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	_, err = chain.WaitContractDeployed(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction mined")
	var sc *ConsumeSkill
	if contractData, err := chain.GetContract(ct.ContractAddr, "Skill"); err != nil {
		t.Fatal(err)
	} else {
		sc = contractData.(*ConsumeSkill)
	}

	events := make(chan *ConsumeSkillStateChange, 128)
	t.Log("watchevent")
	sub, err := chain.watchEvent(context.Background(), sc, events)
	if err != nil {
		t.Fatal(err)
	}
	//call contract.Start
	t.Log("call contract")
	transOpts := bind.NewKeyedTransactor(owner.PrivateKey)
	transOpts.Value = new(big.Int).SetUint64(uint64(contractData.PPrice))
	_, err = sc.Start(transOpts)
	if err != nil {
		t.Fatal(err)
	}
	//mine the contract calling
	simBackend.Commit()
	t.Log("contract call mined")
	//print event
	t.Logf("event received:\n%+v\n", <-events)

	sub.Unsubscribe()
	close(events)
}

func TestConsumeContentContract(t *testing.T) {

	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey)
	t.Logf("ownerAddr:%v", ownerAddr.String())

	alloc := make(core.GenesisAlloc)
	for _, u := range []*user.User{owner} {
		alloc[crypto.PubkeyToAddress(u.PrivateKey.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	}

	simBackend := backends.NewSimulatedBackend(alloc)
	t.Log("sim backend created")
	chain := &ChainEthereum{contractBackend: simBackend, deployBackend: simBackend}
	contractData := struct {
		PPublisher common.Address
		PPlatform  common.Address
		PConsumer  common.Address
		PPrice     uint32
		PRatio     uint8
	}{
		PPublisher: ownerAddr,
		PPlatform:  ownerAddr,
		PConsumer:  ownerAddr,
		PPrice:     1,
		PRatio:     50,
	}
	//start a miner goroutine
	go func() {
		for {
			//mine every seconds
			time.Sleep(time.Second)
			simBackend.Commit()
			t.Log("mined")
		}
	}()
	t.Log("deploy contract")
	ct, err := chain.NewContract(context.Background(), ownerAddr.Bytes(), "Content", contractData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("contract deployment transaction created", common.Bytes2Hex(ct.ContractAddr), "nonce:", ct.RawTx().(*types.Transaction).Nonce())

	//sign the transaction locally, without send private key to the remote
	sig, err := crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction signed by sender")
	err = chain.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction sent to block chain")
	err = chain.WaitMined(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	_, err = chain.WaitContractDeployed(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction mined")
	var sc *ConsumeContent
	if contractData, err := chain.GetContract(ct.ContractAddr, "Content"); err != nil {
		t.Fatal(err)
	} else {
		sc = contractData.(*ConsumeContent)
	}

	//call contract.Start
	t.Log("call contract")
	price, err := sc.Price(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("contract's price:%v", price)
}
func TestConsumeContentContractOnPrivateChain(t *testing.T) {

	keystring, err := ioutil.ReadFile("/Users/dalei/ethereum-nodes/ethereum-node-b/keystore/UTC--2018-04-02T07-45-10.235766815Z--566303d021f916ff6ac743db2514beaadb05b1b6")
	if err != nil {
		t.Fatal("cannot load key from keystore", err)
	}
	key, err := keystore.DecryptKey([]byte(keystring), "123456")
	if err != nil {
		t.Fatal("failed to decrypt key string", err)
	}
	ownerKey := key.PrivateKey
	owner := &user.User{PrivateKey: ownerKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey)
	t.Logf("ownerAddr:%v", ownerAddr.String())

	blc, err := DialEthereum("http://localhost:8545")
	t.Log("sim backend created")
	chain := &ChainEthereum{contractBackend: blc.contractBackend, deployBackend: blc.deployBackend}
	contractData := struct {
		PPublisher []byte
		PPlatform  []byte
		PConsumer  []byte
		PPrice     uint32
		PRatio     uint8
	}{
		PPublisher: ownerAddr.Bytes(),
		PPlatform:  ownerAddr.Bytes(),
		PConsumer:  ownerAddr.Bytes(),
		PPrice:     1,
		PRatio:     50,
	}

	t.Log("deploy contract")
	ct, err := chain.NewContract(context.Background(), ownerAddr.Bytes(), "Content", contractData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("contract deployment transaction created", common.Bytes2Hex(ct.ContractAddr), "nonce:", ct.RawTx().(*types.Transaction).Nonce())

	//sign the transaction locally, without send private key to the remote
	sig, err := crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction signed by sender")
	err = chain.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction sent to block chain")
	err = chain.WaitMined(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("mined")
	_, err = chain.WaitContractDeployed(context.Background(), ct.RawTx())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction deployed")
	var sc *ConsumeContent
	if contractData, err := chain.GetContract(ct.ContractAddr, "Content"); err != nil {
		t.Fatal(err)
	} else {
		sc = contractData.(*ConsumeContent)
	}

	//call contract.Start
	t.Log("call contract")
	_, err = sc.Price(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExtractAbiParams(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(ConsumeContentABI))
	if err != nil {
		t.Fatal(err)
	}
	contractData := struct {
		PHash      string
		PPublisher []byte
		PPlatform  []byte
		PConsumer  []byte
		PPrice     uint32
		PRatio     uint8
	}{
		PHash:      "hash",
		PPublisher: common.Address{}.Bytes(),
		PPlatform:  common.Address{}.Bytes(),
		PConsumer:  common.Address{}.Bytes(),
		PPrice:     10,
		PRatio:     50,
	}
	args, err := extractAbiParams(parsed.Constructor, contractData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", args)
}
