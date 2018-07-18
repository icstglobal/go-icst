package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"reflect"

	"github.com/icstglobal/go-icst/transaction"
)

type Chain interface {
	GetContract(addr []byte, contractType string) (interface{}, error)
	NewContract(ctx context.Context, from []byte, contractType string, contractData interface{}) (*transaction.ContractTransaction, error)
	Call(ctx context.Context, from []byte, contractType string, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.ContractTransaction, error)
	ConfirmTrans(ctx context.Context, trans *transaction.ContractTransaction, sig []byte) error
	WaitMined(ctx context.Context, trans *transaction.ContractTransaction) error
	BalanceAt(ctx context.Context, addr []byte) (*big.Int, error)
	PubKeyToAddress(pub *ecdsa.PublicKey) []byte
	UnmarshalPubkey(pub string) (*ecdsa.PublicKey, error)
	MarshalPubKey(pub *ecdsa.PublicKey) string
	GenerateKey(ctx context.Context) (*ecdsa.PrivateKey, error)
}

// ChainType defines the type of underlying blockchain
type ChainType int

const (
	//Eth blockchain
	Eth ChainType = 0
	//EOS blockchain
	EOS ChainType = 1
)

type ContractType string

const (
	ContentContractType ContractType = "Content"
	SkillContractType   ContractType = "Skill"
)

type ContractEvent struct {
	Addr    []byte       //contract address
	Name    string       //event name
	T       reflect.Type //type of underlying chain specific contract event
	V       interface{}  //underlying chain specific contract event
	Unwatch func()       //unwatch the event at any time
}
