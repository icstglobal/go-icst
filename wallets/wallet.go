package wallets

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/icstglobal/go-icst/chain"
)

//Wallet is the container of accounts.
type Wallet struct {
	ID string
	s  Store
}

//UseAccount selects an account to user
func (w Wallet) UseAccount(ctx context.Context, ID string) (*Account, error) {
	accountBasic, err := w.s.GetAccountBasic(ctx, ID)
	buf, err := base64.StdEncoding.DecodeString(accountBasic.PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key, caused by:%v", err)
	}

	a := &Account{ID: ID, s: w.s}
	var blc chain.Chain
	if blc, err = chain.Get(chain.ChainType(accountBasic.ChainType)); err != nil {
		return nil, err
	}
	a.blc = blc

	var pub crypto.PublicKey
	if pub, err = x509.ParsePKIXPublicKey(buf); err != nil {
		return nil, fmt.Errorf("failed to parse public key, caused by:%v", err)
	}
	if ecdsaPub, ok := pub.(*ecdsa.PublicKey); !ok {
		return nil, fmt.Errorf("not a ecdsa public key, caused by:%v", err)
	} else {
		a.pubkey = ecdsaPub
	}

	return a, nil
}
