package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"reflect"

	"github.com/icstglobal/go-icst/transaction"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/chain/eth"
	"github.com/icstglobal/go-icst/user"
)

var ipc = flag.String("ipc", "", "the geth ipc to attach")
var keyfile = flag.String("keyfile", "", "key file to read private key")
var pwd = flag.String("pwd", "", "password to decrypt the key file")

func main() {
	flag.Parse()
	if len(*ipc) == 0 || len(*keyfile) == 0 || len(*pwd) == 0 {
		flag.PrintDefaults()
		return
	}

	log.SetOutput(os.Stdout)

	keystring, err := ioutil.ReadFile(*keyfile)
	if err != nil {
		log.Fatal("cannot load key from keystore", err)
	}
	key, err := keystore.DecryptKey([]byte(keystring), *pwd)
	if err != nil {
		log.Fatal("failed to decrypt key string", err)
	}
	ownerKey := key.PrivateKey
	owner := &user.User{PrivateKey: ownerKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey)
	log.Printf("ownerAddr:%v", ownerAddr.String())

	blc, err := eth.DialEthereum(*ipc)
	if err != nil {
		log.Fatal("failed to connect to eth", err)
	}
	log.Println("ethereum node connected")
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

	log.Println("deploy contract")
	ct, err := blc.NewContract(context.Background(), ownerAddr.Bytes(), "Content", contractData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("contract deployment transaction created", common.Bytes2Hex(ct.ContractAddr), "nonce:", ct.RawTx().(*types.Transaction).Nonce())

	//sign the transaction locally, without send private key to the remote
	sig, err := crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transaction signed by sender")
	err = blc.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transaction sent to block chain")
	err = blc.WaitMined(context.Background(), ct)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mined")
	_, err = blc.WaitContractDeployed(context.Background(), ct.RawTx())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transaction deployed")

	contractAddr := ct.ContractAddr
	var rawEvt eth.ConsumeContentEventConsume
	log.Println("watchevent")
	events, err := blc.WatchContractEvent(context.Background(), contractAddr, "Content", "EventConsume", reflect.TypeOf(rawEvt))
	if err != nil {
		log.Fatal(err)
	}

	var sc *eth.ConsumeContent

	if ethContract, err := blc.GetContract(contractAddr, "Content"); err != nil {
		log.Fatal(err)
	} else {
		sc = ethContract.(*eth.ConsumeContent)
	}

	transOpts := bind.NewKeyedTransactor(owner.PrivateKey)
	transOpts.Value = new(big.Int).SetUint64(uint64(contractData.PPrice))
	transOpts.GasLimit = 200000
	tx, err := sc.Consume(transOpts)
	ctConsume := transaction.NewContractTransaction(tx, ownerAddr.Bytes())
	err = blc.WaitMined(context.Background(), ctConsume)
	if err != nil {
		log.Fatal(err)
	}
	cnt, err := sc.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("count consumed:%v", cnt.Uint64())

	var e *chain.ContractEvent
	e = <-events
	log.Printf("event:%+v", e)
	e.Unwatch()
}
