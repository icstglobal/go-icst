package wallets

import (
	"context"
	"fmt"

	"github.com/icstglobal/go-icst/chain"
)

var blc chain.Chain

//Wallet is the container of accounts.
type Wallet struct {
	ID string
	s  Store
}

func NewWallet(walletID string, s Store) *Wallet {
	return &Wallet{ID: walletID, s: s}
}

//ImportAccount save an account to user
func (w *Wallet) ImportAccount(ctx context.Context, pubKey string, chainType chain.ChainType) (*AccountRecordBasic, error) {
	return w.s.SetAccountBasic(ctx, w.ID, pubKey, chainType)
}

func (w *Wallet) ExistAccount(ctx context.Context, pubKey string, chainType chain.ChainType) bool {
	return w.s.ExistAccount(ctx, pubKey, chainType)
}

func (w *Wallet) GetAccounts(ctx context.Context) ([]*AccountRecordBasic, error) {
	return w.s.GetAccounts(ctx, w.ID)
}

func (w *Wallet) GetAccountBasic(ctx context.Context, accountId string) (*AccountRecordBasic, error) {
	return w.s.GetAccountBasic(ctx, accountId)
}

//UseAccount selects an account to user
func (w *Wallet) UseAccount(ctx context.Context, accountID string) (*Account, error) {
	accountBasic, err := w.s.GetAccountBasic(ctx, accountID)

	a := &Account{ID: accountBasic.ID, s: w.s}
	var blc chain.Chain
	if blc, err = chain.Get(chain.ChainType(accountBasic.ChainType)); err != nil {
		return nil, err
	}
	a.blc = blc
	a.PublicKey, err = blc.UnmarshalPubkey(accountBasic.PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key, caused by:%v", err)
	}

	return a, nil
}
