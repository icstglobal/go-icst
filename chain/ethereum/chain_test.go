package ethereum

import (
	"context"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/icstglobal/go-icst/skill"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/icstglobal/go-icst/contract"
	"github.com/icstglobal/go-icst/user"
)

func TestConsumeSkillContract(t *testing.T) {

	ownerKey, _ := crypto.GenerateKey()
	owner := &user.User{PrivateKey: ownerKey}

	alloc := make(core.GenesisAlloc)
	for _, u := range []*user.User{owner} {
		alloc[crypto.PubkeyToAddress(u.PrivateKey.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	}

	simBackend := backends.NewSimulatedBackend(alloc)
	t.Log("sim backend created")
	chain := &ChainEthereum{contractBackend: simBackend, deployBackend: simBackend}
	skill := &contract.SkillContract{
		Skill:    &skill.Skill{Hash: []byte("hex"), Data: []byte("test"), Producer: owner},
		Options:  &contract.Options{Platform: owner, Price: 1, Ratio: 50},
		Consumer: owner,
		Price:    1,
	}
	//start a miner goroutine
	go func() {
		//mine after 2 seconds
		time.Sleep(2 * time.Second)
		simBackend.Commit()
		t.Log("contract deployment mined")
	}()
	t.Log("deploy contract")
	addr, err := chain.DeployContract(context.Background(), skill)
	if err != nil {
		t.Fatal(err)
	}
	var sc *ConsumeSkill
	if icontract, err := chain.GetContract(addr, reflect.TypeOf(skill)); err != nil {
		t.Fatal(err)
	} else {
		sc = icontract.(*ConsumeSkill)
	}

	events := make(chan *ConsumeSkillStateChange, 128)
	t.Log("watchevent")
	sub, err := chain.WatchEvent(context.Background(), sc, events)
	if err != nil {
		t.Fatal(err)
	}
	//call contract.Start
	t.Log("call contract")
	transOpts := bind.NewKeyedTransactor(owner.PrivateKey)
	transOpts.Value = new(big.Int).SetUint64(uint64(skill.Price))
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
