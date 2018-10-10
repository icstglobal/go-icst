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
	NewContract(ctx context.Context, from []byte, contractType string, contractData interface{}) (*transaction.Transaction, error)
	Call(ctx context.Context, from []byte, contractType string, contractAddr []byte, methodName string, value *big.Int, callData interface{}) (*transaction.Transaction, error)
	CallWithAbi(ctx context.Context, from []byte, contractAddr []byte, methodName string, value *big.Int, callData interface{}, abiStr string) (*transaction.Transaction, error)
	ConfirmTrans(ctx context.Context, trans *transaction.Transaction, sig []byte) error
	WaitMined(ctx context.Context, trans *transaction.Transaction) error
	BalanceAt(ctx context.Context, addr []byte) (*big.Int, error)
	BalanceAtICST(ctx context.Context, addr []byte) (*big.Int, error)
	PubKeyToAddress(pub *ecdsa.PublicKey) []byte
	UnmarshalPubkey(pub string) (*ecdsa.PublicKey, error)
	MarshalPubKey(pub *ecdsa.PublicKey) string
	GenerateKey(ctx context.Context) (*ecdsa.PrivateKey, error)
	Sign(hash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error)
	Transfer(ctx context.Context, from []byte, to []byte, value *big.Int) (*transaction.Transaction, error)
	TransferICST(ctx context.Context, from []byte, to []byte, value *big.Int) (*transaction.Transaction, error)
	WatchBlocks(ctx context.Context, blockStart *big.Int) (<-chan *transaction.Block, <-chan error)
	WatchICSTTransfer(ctx context.Context, blockStart *big.Int) (<-chan *transaction.Block, <-chan error)

	// GetEvents(ctx context.Context, topics [][]common.Hash, fromBlock *big.Int) ([]interface{}, error)
	GetContractEvents(ctx context.Context, addr []byte, fromBlock, toBlock *big.Int, abiString string, eventName string, eventVType reflect.Type) ([]*ContractEvent, error)
	UnpackLog(abiStr string, out interface{}, event string, log interface{}) error
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
	Addr      []byte       //contract address
	Name      string       //event name
	T         reflect.Type //type of underlying chain specific contract event
	V         interface{}  //underlying chain specific contract event
	BlockNum  uint64
	BlockHash []byte
	TxIndex   uint64
	TxHash    []byte
	Unwatch   func() //unwatch the event at any time
}
