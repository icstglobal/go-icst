package wallets

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/icstglobal/go-icst/common"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/transaction"
)

//Account maps the block chain account, which has a public address and holds tokens
type Account struct {
	ID        string
	PublicKey *ecdsa.PublicKey
	blc       chain.Chain
	s         Store
}

//GetBalance returns the current balance of this account
func (a *Account) GetBalance(ctx context.Context) (*big.Int, error) {
	return a.blc.BalanceAt(ctx, a.Addr())
}

//BackupKey stores the ecypted key file, so that the user can recover it later
func (a *Account) BackupKey(ctx context.Context, encryptedKey string, r *big.Int, s *big.Int, hint string, encryptedHint string) error {
	//check the signature of the cypher text and store it
	//The cypher text must be signed by the account's public key

	encryptedKeyBytes, err := hex.DecodeString(encryptedKey)
	if err != nil {
		return err
	}

	hash := common.Hash(encryptedKeyBytes)
	if valid := ecdsa.Verify(a.PublicKey, hash[:], r, s); !valid {
		return common.ErrorInvalidSignature
	}
	return a.s.SaveKey(ctx, a.ID, encryptedKey, hint, encryptedHint)
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
	return a.blc.PubKeyToAddress(a.PublicKey)
}

// CreateContentContractTrans creates contract transaction
func (a *Account) CreateContentContractTrans(ctx context.Context, data map[string]interface{}) (*transaction.Transaction, error) {
	ownerAddr := a.Addr()
	publisher := content.NewPublisher(a.blc, nil)
	trans, err := publisher.Pub(context.Background(), ownerAddr, data)
	if err != nil {
		return nil, err
	}

	return trans, nil
}

// AfterSign process after client sign
// including confirm transaction and wait for mining
func (a *Account) AfterSign(ctx context.Context, sigHex string, trans *transaction.Transaction) error {
	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return err
	}

	err = a.blc.ConfirmTrans(context.Background(), trans, sig)
	if err != nil {
		return fmt.Errorf("failed to confirm contract creation transaction:%v", err)
	}
	if err = a.blc.WaitMined(context.Background(), trans); err != nil {
		return fmt.Errorf("error happen when wait transaction mined:%v", err)
	}
	return nil
}

// Create Contract Transaction
func (a *Account) CallContentContract(ctx context.Context, cxAddrStr string, data map[string]interface{}) (*transaction.Transaction, error) {
	ownerAddr := a.Addr()
	method := data["Method"].(string)
	price := data["Price"].(int)
	callData := data["CallData"]

	cxAddr, err := hex.DecodeString(cxAddrStr)
	if err != nil {
		return nil, err
	}
	trans, err := a.blc.Call(context.TODO(), ownerAddr, "Content", cxAddr, method, big.NewInt(int64(price)), callData)
	if err != nil {
		return nil, err
	}

	return trans, nil
}

func (a *Account) GetBlc() chain.Chain {
	return a.blc
}

//Transfer generate the transaction to transfer
func (a *Account) Transfer(ctx context.Context, to []byte, val *big.Int) (*transaction.Transaction, error) {
	return a.blc.Transfer(ctx, a.Addr(), to, val)
}

//TransferICST generate the transaction to transfer ICST
func (a *Account) TransferICST(ctx context.Context, to []byte, val *big.Int) (*transaction.Transaction, error) {
	return a.blc.TransferICST(ctx, a.Addr(), to, val)
}
