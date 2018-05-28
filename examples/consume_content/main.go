package main

import (
	"context"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/chain/eth"
	"github.com/icstglobal/go-icst/contract"

	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/user"
)

var ethURL = flag.String("ethURL", "http://localhost:8545", "the endpoint of a Ethereum rpc service")
var contractAddrString = flag.String("addrHex", "", "the address of smart contract deployed already, in Hex without '0x'")
var doConsume = flag.Bool("doConsume", false, "consume the content ot not, default to false")
var ownerKeyFile = flag.String("ownerKeyFile", "", "owner's key file path")
var ownerKeyPwd = flag.String("ownerKeyPwd", "", "password for owner's key")
var platformKeyFile = flag.String("platformKeyFile", "", "platform's key file path")
var platFormKeyPwd = flag.String("platformKeyPwd", "", "password for platform's key")
var consumerKeyFile = flag.String("consumerKeyFile", "", "consumer's key file path")
var consumerKeyPwd = flag.String("consumerKeyPwd", "", "password for consumer's key")

var ownerKey *keystore.Key
var platformKey *keystore.Key
var consumerKey *keystore.Key

func main() {
	flag.Parse()
	initAccounts()

	chain, err := ethChain(*ethURL)
	if err != nil {
		log.Fatal(err)
	}

	testData := "test data"
	owner := &user.User{PrivateKey: ownerKey.PrivateKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey).Bytes()
	platform := &user.User{PrivateKey: platformKey.PrivateKey}
	platformAddr := crypto.PubkeyToAddress(platform.PrivateKey.PublicKey).Bytes()
	consumer := &user.User{PrivateKey: consumerKey.PrivateKey}
	opts := contract.Options{Platform: platformAddr, Price: 1, Ratio: 50}

	var addr []byte
	// If contract address not given, deploy a new one
	if contractAddrString == nil || *contractAddrString == "" {
		publisher := content.NewPublisher(chain, nil)
		//contract options
		contractData := make(map[string]interface{})
		contractData["PPublisher"] = ownerAddr
		contractData["PPlatform"] = opts.Platform
		contractData["PHash"] = []byte(testData)
		contractData["PPrice"] = opts.Price
		contractData["PRatio"] = opts.Ratio

		trans, err := publisher.Pub(context.Background(), ownerAddr, contractData)
		if err != nil {
			log.Fatal(err)
		}
		sig, err := crypto.Sign(trans.Hash(), owner.PrivateKey)
		if err != nil {
			log.Fatal("failed to sign a transaction", err)
		}
		err = chain.ConfirmTrans(context.Background(), trans, sig)
		if err != nil {
			log.Fatal("failed to confirm contract creation transaction")
		}
		if err = chain.WaitMined(context.Background(), trans.RawTx()); err != nil {
			log.Fatal("error happen when wait transaction mined", err)
		}
		log.Printf("smart contract deployed to address:%v\n", common.BytesToHash(addr).Hex())
	} else {
		addr, err = hex.DecodeString(*contractAddrString)
		if err != nil {
			log.Fatal("contract address given is not a valid hex,", err)
		}
	}

	ct, err := chain.GetContract(addr, "Content")
	if err != nil {
		log.Fatal(err)
	}

	contentContract := ct.(*eth.ConsumeContent)
	cnt, err := contentContract.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("count before consumed:%v", cnt)

	if !*doConsume {
		log.Print("consume is not needed")
		return
	}

	callData := new(struct{})
	trans, err := chain.Call(context.TODO(), ownerAddr, "Content", addr, "consume", big.NewInt(int64(opts.Price)), callData)
	sig, err := crypto.Sign(trans.Hash(), consumer.PrivateKey)
	if err != nil {
		log.Fatal("failed to sign a transaction", err)
	}
	err = chain.ConfirmTrans(context.Background(), trans, sig)
	if err != nil {
		log.Fatal("failed to confirm contract creation transaction")
	}
	if err = chain.WaitMined(context.Background(), trans.RawTx()); err != nil {
		log.Fatal("error happen when wait transaction mined", err)
	}

	cnt, err = contentContract.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("count after consumed:%v", cnt)
}

func ethChain(url string) (chain.Chain, error) {
	// client, err := rpc.DialHTTP(url)

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("failed to connect eth rpc endpoint {%v}, err is:%v \n", url, err)
		return nil, err
	}

	return eth.NewChainEthereum(client), nil
}

func initAccounts() {
	ownerKeyString, err := ioutil.ReadFile(*ownerKeyFile)
	if err != nil {
		log.Fatal(err, *ownerKeyFile)
	}
	ownerKey, err = keystore.DecryptKey(ownerKeyString, *ownerKeyPwd)
	if err != nil {
		log.Fatal(err)
	}

	platformKeyString, err := ioutil.ReadFile(*platformKeyFile)
	if err != nil {
		log.Fatal(err, *platformKeyFile)
	}
	platformKey, err = keystore.DecryptKey(platformKeyString, *platFormKeyPwd)
	if err != nil {
		log.Fatal(err)
	}

	consumerKeyString, err := ioutil.ReadFile(*consumerKeyFile)
	if err != nil {
		log.Fatal(err, *consumerKeyFile)
	}
	consumerKey, err = keystore.DecryptKey(consumerKeyString, *consumerKeyPwd)
	if err != nil {
		log.Fatal(err)
	}
}
