package chain

import (
	"errors"
	"fmt"

	"github.com/icstglobal/go-icst/chain/ethereum"
	"github.com/icstglobal/go-icst/contract"
)

type Chain interface {
	DeployContract(contract *contract.ConsumeContract) (contractAddr []byte, err error)
	GetContract(addr []byte) (interface{}, error)
}

// ChainType defines the type of underlying blockchain
type ChainType int

const (
	//Ethereum blockchain
	Ethereum ChainType = 0
	//EOS blockchain
	EOS ChainType = 1
)

//NewChain inits a chain instance by chain type
func NewChain(t ChainType) (Chain, error) {
	switch t {
	case Ethereum:
		chain := &ethereum.ChainEthereum{}
		return chain, nil
	case EOS:
		return nil, nil
	default:
		msg := fmt.Sprintf("unsuported chain type:%v", t)
		return nil, errors.New(msg)
	}
}
