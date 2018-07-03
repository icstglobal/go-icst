package wallets

import (
	"context"
	"github.com/icstglobal/go-icst/chain"
	db "github.com/icstglobal/go-icst/wallets/database"
	"log"
	"github.com/pborman/uuid"
)

type Store interface {
	GetKeyHint(ctx context.Context, accountID string) (string, error)
	GetKey(ctx context.Context, accountID string, encryptedHint string) (string, error)
	SaveKey(ctx context.Context, accountID string, encryptedKey string) error
	GetAccountBasic(ctx context.Context, accountID string) (AccountRecordBasic, error)
	GetAccounts(ctx context.Context, walletID string) ([]AccountRecordBasic, error)
	SetAccountBasic(ctx context.Context, walletID string, pubKey string, chainType chain.ChainType) (AccountRecordBasic, error)
	IsExistAccount(ctx context.Context, pubKey string, chainType chain.ChainType) bool
	HasAccount(ctx context.Context, accountID string) bool
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

func (a *AccountRecord) GetAccounts(ctx context.Context, walletID string) ([]AccountRecordBasic, error) {
	var accounts = []AccountRecordBasic{}
	var (
		accountID string
		pubKey string
		chainType int
		createTime string
	)
	rows, err := db.DBCon.Query("select accountID, pubKey, chainType, createTime from account where walletID = ?", walletID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&accountID, &pubKey, &chainType, &createTime)
		if err != nil {
			log.Fatal(err)
		}
		account := AccountRecordBasic{ID:accountID, ChainType:chain.ChainType(chainType), PubKey: pubKey}
		accounts = append(accounts, account)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return accounts, nil
}


func (a *AccountRecord) GetAccountBasic(ctx context.Context, accountID string) (AccountRecordBasic, error) {
					
	var chainType int
	var pubKey string
	err := db.DBCon.QueryRow("SELECT chainType, PubKey FROM account WHERE accountID=?", accountID).Scan(&chainType, &pubKey)
	if err != nil {
		log.Fatal(err)
	}	
	// pubKey = "BLA8u6oCXN5qSf0pzQ5UDcDttiil2T5VR52ie5lpLG3e1RsBbkAoibtsHFYv3qe5yA0qzZWDNYWLaLrwnIHQjWo="
	return AccountRecordBasic{ID:accountID, ChainType:chain.ChainType(chainType), PubKey: pubKey}, nil
}


func (a *AccountRecord) HasAccount(ctx context.Context, accountID string) bool {
	var count int
	err := db.DBCon.QueryRow("SELECT count(accountID) as count FROM account WHERE accountID=?", accountID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}	
	if count == 1{
		return true
	}
	return false
}

func (a *AccountRecord) IsExistAccount(ctx context.Context, pubKey string, chainType chain.ChainType) bool {
	var count int
	err := db.DBCon.QueryRow("SELECT count(accountID) as count FROM account WHERE pubKey=? and chainType=?", pubKey, chainType).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}	
	if count == 1{
		return true
	}
	return false
}

func (a *AccountRecord) SetAccountBasic(ctx context.Context, walletID string, pubKey string, chainType chain.ChainType) (AccountRecordBasic, error){
	accountID := uuid.New() 
	stmt, err := db.DBCon.Prepare("INSERT INTO account(accountID, walletID, pubKey, chainType) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(accountID, walletID, pubKey, chainType)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	account := AccountRecordBasic{ID:accountID, ChainType:chain.ChainType(chainType), PubKey: pubKey}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return account, nil
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
