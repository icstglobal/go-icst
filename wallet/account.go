package wallet

import (
	"context"
	"math/big"

	"github.com/icstglobal/go-icst/chain"
)

//Account maps the block chain account, which has a public address and holds tokens
type Account struct {
	addr []byte
	blc  chain.Chain
}

//UseAddress selects an account at the given address
func UseAddress(t chain.ChainType, addr []byte) (*Account, error) {
	a := &Account{addr: addr}
	if blc, err := chain.Get(t); err != nil {
		return nil, err
	} else {
		a.blc = blc
		return a, nil
	}
}

//GetBalance returns the current balance of this account
func (a *Account) GetBalance(ctx context.Context) (*big.Int, error) {
	return a.blc.BalanceAt(ctx, a.addr)
}
