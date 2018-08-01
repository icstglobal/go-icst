package transaction

import (
	"math/big"
)

type Block struct {
	BlockNumber uint64
	Hash        string
	Trans       []*Message
}

//Message is the immutable transaction
type Message struct {
	Hash       string
	To         []byte
	From       []byte
	Nonce      uint64
	Amount     *big.Int
	GasLimit   uint64
	GasPrice   *big.Int
	Data       []byte
	CheckNonce bool
}
