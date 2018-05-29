package wallets

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/icstglobal/go-icst/common"

	"github.com/icstglobal/go-icst/chain"
)

//Account maps the block chain account, which has a public address and holds tokens
type Account struct {
	ID     string
	pubkey *ecdsa.PublicKey
	blc    chain.Chain
	s      Store
}

//GetBalance returns the current balance of this account
func (a *Account) GetBalance(ctx context.Context) (*big.Int, error) {
	return a.blc.BalanceAt(ctx, a.Addr())
}

//BackupKey stores the ecypted key file, so that the user can recover it later
func (a *Account) BackupKey(ctx context.Context, encryptedKey string, r *big.Int, s *big.Int) error {
	//check the signature of the cypher text and store it
	//The cypher text must be signed by the account's public key
	hash := common.Hash([]byte(encryptedKey))
	if valid := ecdsa.Verify(a.pubkey, hash[:], r, s); !valid {
		return common.ErrorInvalidSignature
	}
	return a.s.SaveKey(ctx, a.ID, encryptedKey)
}

//RecoverKey returns the encrypted key file to the user. Don't do any decryption.
//The check is the key hint encrypted by user's passphase, and it needs to be matched
func (a *Account) RecoverKey(ctx context.Context, encryptedHint string) (string, error) {
	key, err := a.s.GetKey(ctx, a.ID, encryptedHint)
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return "", errors.New("check for key recovering is invalid")
	}

	return key, nil
}

//RecoverKeyHint returns the hint to recover the account's key.
//The hint will be sent back to user, encrypted by user's passphase and then sent to server for key recovering.
func (a *Account) RecoverKeyHint(ctx context.Context) (string, error) {
	return a.s.GetKeyHint(ctx, a.ID)
}

//Addr returns the address of the account
func (a *Account) Addr() []byte {
	return a.blc.PubKeyToAddress(a.pubkey)
}
