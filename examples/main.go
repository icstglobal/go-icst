package main

import (
	"log"
	"math/big"

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
	testData := "test content"
	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}
	content := &content.Content{Owner: owner, Data: []byte(testData)}
	platformKey, _ := crypto.GenerateKey()
	platform := &user.User{PrivateKey: platformKey}
	//test chain
	simChain, _ := SimChain([]*user.User{owner, platform})
	publisher := publish.NewContentPublisher(simChain)
	//contract options
	opts := contract.Options{Platform: platform, Price: 1, Ratio: 50}
	var err error
	addr, err := publisher.PubContent(content, opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("smart contract deployed to address:%v\n", common.BytesToHash(addr).Hex())
}

// SimChain creates a local simulated blockchain, and init balance for users
func SimChain(accounts []*user.User) (chain.Chain, error) {
	alloc := make(core.GenesisAlloc)
	for _, u := range accounts {
		alloc[crypto.PubkeyToAddress(u.PrivateKey.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	}

	return &ethereum.ChainEthereum{Backend: backends.NewSimulatedBackend(alloc)}, nil
}
