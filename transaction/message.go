package transaction

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type Block struct {
	BlockNumber uint64     `json:"blockNumer"`
	Hash        string     `json:"hash"`
	Trans       []*Message `json:"trans"`
}

//Message is the immutable transaction
type Message struct {
	Hash       string   `json:"hash"`
	To         []byte   `json:"to"`
	From       []byte   `json:"from"`
	Nonce      uint64   `json:"nonce"`
	Amount     *big.Int `json:"amount"`
	GasLimit   uint64   `json:"gasLimit"`
	GasPrice   *big.Int `json:"gasPrice"`
	Data       []byte   `json:"data"`
	CheckNonce bool     `json:"checkNonce"`
	Success    bool     `json:"success"`
	GasUsed    uint64   `json:"gasUsed"`
}

func (m Message) String() string {
	return fmt.Sprintf("Hash:%v From:%v To:%v Amount:%v Nonce:%v Success:%v GasUsed:%v GasLimit:%v GasPrice:%v CheckNonce:%v Data:%v",
		m.Hash, hex.EncodeToString(m.From), hex.EncodeToString(m.To), m.Amount, m.Nonce, m.Success, m.GasUsed, m.GasLimit, m.GasPrice,
		m.CheckNonce, m.Data)
}
