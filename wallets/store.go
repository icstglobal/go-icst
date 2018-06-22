package wallets

import (
	"context"
	"github.com/icstglobal/go-icst/chain"
)

type Store interface {
	GetKeyHint(ctx context.Context, accountID string) (string, error)
	GetKey(ctx context.Context, accountID string, encryptedHint string) (string, error)
	SaveKey(ctx context.Context, accountID string, encryptedKey string) error
	GetAccountBasic(ctx context.Context, accountID string) (AccountRecordBasic, error)
	SetAccountBasic(ctx context.Context, accountID string, pubKey string, chainType chain.ChainType) error
	IsExistAccount(ctx context.Context, accountID string) bool
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
}

//AccountRecordSec is the security part of an account
type AccountRecordSec struct {
	EncryptedPrivKey string //encrypted, base 64
	Hint             string //hint to recover key
	EncryptedHint    string
}


func (a *AccountRecord) GetAccountBasic(ctx context.Context, accountID string) (AccountRecordBasic, error) {
	PubKey := "BAqvsVcQ1Qhuj/hZ3nuds/GQqdLkt4cVHGFxmRZh0G24vjs1dRQXJAZgfR4ZCGVB99Dfi4C2GnAU0lEEMi+EjdQ="
	return AccountRecordBasic{ID:accountID, ChainType:chain.Eth, PubKey: PubKey}, nil
}

func (a *AccountRecord) IsExistAccount(ctx context.Context, accountID string) bool {
	return true
}

func (a *AccountRecord) SetAccountBasic(ctx context.Context, accountID string, pubKey string, chainType chain.ChainType) error{
	return nil
}

func (a *AccountRecord) GetKeyHint(ctx context.Context, accountID string) (string, error) {
	return "", nil
}

func (a *AccountRecord) GetKey(ctx context.Context, accountID string, encryptedHint string) (string, error) {
	return "", nil
}


func (a *AccountRecord) SaveKey(ctx context.Context, accountID string, encryptedKey string) error{
	return nil
}
