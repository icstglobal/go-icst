package wallets

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"encoding/hex"
	"fmt"

	"github.com/icstglobal/go-icst/common"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/transaction"
	"github.com/icstglobal/go-icst/content"
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


// Create Contract Transaction
func (a *Account) CreateContentContractTrans(ctx context.Context, data map[string]interface{}) (*transaction.ContractTransaction, error) {
	ownerAddr := a.Addr()
	publisher := content.NewPublisher(a.blc, nil)
	trans, err := publisher.Pub(context.Background(), ownerAddr, data)
	if err != nil{
		return nil, err
	}

	return trans, nil
}

// process after client sign
// including confirm transaction and wait for mining
func (a *Account) AfterSign(ctx context.Context, sigHex string, trans *transaction.ContractTransaction) (error) {
	sig, err := hex.DecodeString(sigHex)
	if err != nil{
		return err
	}

	err = a.blc.ConfirmTrans(context.Background(), trans, sig)
	if err != nil {
		return fmt.Errorf("failed to confirm contract creation transaction", err)
	}
	if err = a.blc.WaitMined(context.Background(), trans); err != nil {
		return fmt.Errorf("error happen when wait transaction mined", err)
	}
	return nil
}
