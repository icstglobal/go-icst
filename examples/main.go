package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/chain/ethereum"
	"github.com/icstglobal/go-icst/contract"
	"github.com/icstglobal/go-icst/publish"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/user"
)

func main() {
	//log to console
	log.SetOutput(os.Stdout)

	testData := "test content"
	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}
	content := &content.Content{Owner: owner, Data: []byte(testData)}
	platformKey, _ := crypto.GenerateKey()
	platform := &user.User{PrivateKey: platformKey}
	//test chain
	simChain, simBackend, _ := SimChain([]*user.User{owner, platform})
	publisher := publish.NewContentPublisher(simChain)
	//contract options
	opts := contract.Options{Platform: platform, Price: 1, Ratio: 50}
	var err error
	addr, err := publisher.PubContent(content, opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("smart contract deployed to address:%v\n", common.BytesToHash(addr).Hex())

	fmt.Println("mining")
	simBackend.Commit()

	ct, err := simChain.GetContract(addr)
	if err != nil {
		log.Fatal(err)
	}

	contentContract := ct.(*ethereum.ConsumeContent)
	cnt, err := contentContract.Count(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("count:%v\n", cnt) //output:0

	transactor := bind.NewKeyedTransactor(owner.PrivateKey)
	transactor.Value = big.NewInt(int64(opts.Price))
	//we dont care about the transaction now
	_, err = contentContract.Consume(transactor)
	if err != nil {
		log.Fatal(err)
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
	return &ethereum.ChainEthereum{Backend: simBackend}, simBackend, nil
}
