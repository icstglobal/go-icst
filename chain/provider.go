package chain

import (
	"context"
	"reflect"

	"github.com/icstglobal/go-icst/chain/ethereum"
)

type Chain interface {
	GetContract(addr []byte, t reflect.Type) (interface{}, error)
	DeployContract(ctx context.Context, icontract interface{}) (contractAddr []byte, err error)
}

// ChainType defines the type of underlying blockchain
type ChainType int

const (
	//Ethereum blockchain
	Ethereum ChainType = 0
	//EOS blockchain
	EOS ChainType = 1
)

func DialEthereum(url string) (Chain, error) {
	return ethereum.DialEthereum(url)
}
