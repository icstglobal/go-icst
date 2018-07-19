package wallets

import (
	"context"

	"github.com/icstglobal/go-icst/chain"
)

type Store interface {
	GetKeyHint(ctx context.Context, accountID string) (string, error)
	GetKey(ctx context.Context, accountID string, encryptedHint string) (string, error)
	SaveKey(ctx context.Context, accountID string, encryptedKey string, hint string, encryptedHint string) error
	GetAccountBasic(ctx context.Context, accountID string) (*AccountRecordBasic, error)
	GetAccounts(ctx context.Context, walletID string) ([]*AccountRecordBasic, error)
	SetAccountBasic(ctx context.Context, walletID string, pubKey string, address string, chainType chain.ChainType) (*AccountRecordBasic, error)
	ExistAccount(ctx context.Context, pubKey string, chainType chain.ChainType) bool
}

type AccountRecord struct {
	AccountRecordBasic
	AccountRecordSec
}

//AccountRecordBasic is the basic part of an account
type AccountRecordBasic struct {
	ID        string
	ChainType chain.ChainType
	PubKey    string //base 64
	Address   string
}

//AccountRecordSec is the security part of an account
type AccountRecordSec struct {
	EncryptedPrivKey string //encrypted, base 64
	Hint             string //hint to recover key
	EncryptedHint    string
}
