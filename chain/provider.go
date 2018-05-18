package chain

import (
	"context"
	"math/big"

	"github.com/icstglobal/go-icst/transaction"

	"github.com/icstglobal/go-icst/chain/ethereum"
)

type Chain interface {
	GetContract(addr []byte, contractType string) (interface{}, error)
	NewContract(ctx context.Context, from []byte, contractType string, contractData interface{}) (*transaction.ContractTransaction, error)
	Call(ctx context.Context, from []byte, contractType string, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.ContractTransaction, error)
	ConfirmTrans(ctx context.Context, trans *transaction.ContractTransaction, sig []byte) error
	WaitMined(ctx context.Context, tx interface{}) error
}

// ChainType defines the type of underlying blockchain
type ChainType int

const (
	//Ethereum blockchain
	Ethereum ChainType = 0
	//EOS blockchain
	EOS ChainType = 1
)

type ContractType string

const (
	ContentContractType ContractType = "Content"
	SkillContractType   ContractType = "Skill"
)

func DialEthereum(url string) (Chain, error) {
	return ethereum.DialEthereum(url)
}
