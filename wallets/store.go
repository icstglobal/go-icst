package wallets

import (
	"context"
)

type Store interface {
	GetKeyHint(ctx context.Context, accountID string) (string, error)
	GetKey(ctx context.Context, accountID string, encryptedHint string) (string, error)
	SaveKey(ctx context.Context, accountID string, encryptedKey string) error
	GetAccountBasic(ctx context.Context, accountID string) (AccountRecordBasic, error)
}

type AccountRecord struct {
	AccountRecordBasic
	AccountRecordSec
}

//AccountRecordBasic is the basic part of an account
type AccountRecordBasic struct {
	ID        string
	ChainType int
	PubKey    string //base 64
}

//AccountRecordSec is the security part of an account
type AccountRecordSec struct {
	EncryptedPrivKey string //encrypted, base 64
	Hint             string //hint to recover key
	EncryptedHint    string
}
