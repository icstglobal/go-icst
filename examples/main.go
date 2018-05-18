package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/chain/ethereum"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/contract"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/user"
)

func main() {
	//log to console
	log.SetOutput(os.Stdout)

	testData := "test content"
	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}
	ownerAddr := crypto.PubkeyToAddress(owner.PrivateKey.PublicKey)
	platformKey, _ := crypto.GenerateKey()
	platform := &user.User{PrivateKey: platformKey}
	platformAddr := crypto.PubkeyToAddress(platform.PrivateKey.PublicKey)
	//test chain
	simChain, simBackend, _ := SimChain([]*user.User{owner, platform})
	publisher := content.NewPublisher(simChain, nil)
	//contract options
	opts := contract.Options{Platform: platformAddr.Bytes(), Price: 1, Ratio: 50}
	var err error
	contractData := make(map[string]interface{})
	contractData["PPublisher"] = ownerAddr.Bytes()
	contractData["PPlatform"] = opts.Platform
	contractData["PHash"] = testData
	contractData["PPrice"] = opts.Price
	contractData["PRatio"] = opts.Ratio
	ct, err := publisher.Pub(context.Background(), ownerAddr.Bytes(), contractData)
	if err != nil {
		log.Fatal(err)
	}
	sig, err := crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = simChain.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mining")
	simBackend.Commit()

	simChain.WaitMined(context.Background(), ct.RawTx())
	log.Printf("smart contract deployed to address:%v\n", ct.ContractAddr)

	ctr, err := simChain.GetContract(ct.ContractAddr, string(chain.ContentContractType))
	if err != nil {
		log.Fatal(err)
	}

	contentContract := ctr.(*ethereum.ConsumeContent)
	cnt, err := contentContract.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("count:%v\n", cnt) //output:0

	callData := new(struct{})
	ct, err = simChain.Call(context.TODO(), ownerAddr.Bytes(), "Content", ct.ContractAddr, "consume", big.NewInt(int64(opts.Price)), callData)
	sig, err = crypto.Sign(ct.Hash(), owner.PrivateKey)
	if err != nil {
		log.Fatal("failed to sign a transaction", err)
	}
	err = simChain.ConfirmTrans(context.Background(), ct, sig)
	if err != nil {
		log.Fatal("failed to confirm contract creation transaction")
	}

	fmt.Println("mining")
	simBackend.Commit()

	//count should be increased by 1 because of the consuming
	cnt, err = contentContract.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("count:%v\n", cnt) //output:1
}

// SimChain creates a local simulated blockchain, and init balance for users
func SimChain(accounts []*user.User) (chain.Chain, *backends.SimulatedBackend, error) {
	alloc := make(core.GenesisAlloc)
	for _, u := range accounts {
		alloc[crypto.PubkeyToAddress(u.PrivateKey.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	}

	simBackend := backends.NewSimulatedBackend(alloc)
	return ethereum.NewSimChainEthereum(simBackend), simBackend, nil
}
