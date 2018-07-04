package wallets

import (
	"context"
	"fmt"

	"github.com/icstglobal/go-icst/chain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/icstglobal/go-icst/chain/eth"
	db "github.com/icstglobal/go-icst/wallets/database"
	conf "github.com/icstglobal/go-icst/config"
)

var blc chain.Chain


//Wallet is the container of accounts.
type Wallet struct {
	ID string
	s  Store
}

type ChainConf struct{
	ChainType int
	Url string
}

func NewWallet(walletID string) *Wallet {
	return &Wallet{ID:walletID, s:&AccountRecord{}}
}


// init env
// 1. dail main chain, init chain interface and save to pool
// 2. init wallet store
func Init(chainConfs []*ChainConf, confFile string) (error){
    // init db
    conf.Load(confFile)
    // conf.Config.Mysql
    db.DBCon = db.DB()

    // init chain
	for _, chainConf:= range(chainConfs){
		if chainConf.ChainType == int(chain.Eth){
			//dial eth chain
			url := chainConf.Url
			client, err := ethclient.Dial(url)
			if err != nil {
				return fmt.Errorf("failed to connect eth rpc endpoint {%v}, err is:%v \n", url, err)
			}
			blc = eth.NewChainEthereum(client)
			chain.Set(chain.Eth, blc)
		}
		if chainConf.ChainType == int(chain.EOS){
			fmt.Println("not support EOS yet.")
		}
	}

	return nil
}

//SetAccount save an account to user
func (w *Wallet) SetAccount(ctx context.Context, walletID string, pubKey string, chainType chain.ChainType) (AccountRecordBasic, error) {
	return w.s.SetAccountBasic(ctx, walletID, pubKey, chainType)
}


func (w *Wallet) IsExistAccount(ctx context.Context, pubKey string, chainType chain.ChainType) bool{
	return w.s.IsExistAccount(ctx, pubKey, chainType)
}

func (w *Wallet) HasAccount(ctx context.Context, accountID string) bool{
	return w.s.HasAccount(ctx, accountID)
}

func (w *Wallet) GetAccounts(ctx context.Context, walletID string) ([]AccountRecordBasic, error){
	return w.s.GetAccounts(ctx, walletID)
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
	a.pubkey, err = blc.UnmarshalPubkey(accountBasic.PubKey)
	if err != nil{
		return nil, fmt.Errorf("failed to parse public key, caused by:%v", err)
	}

	return a, nil
}

