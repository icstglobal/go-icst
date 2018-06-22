package wallets

import (
	"context"
	// "crypto"
	// "crypto/ecdsa"
	// "crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/transaction"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/icstglobal/go-icst/chain/eth"
	gcrypto "github.com/ethereum/go-ethereum/crypto"
)

var blc chain.Chain


//Wallet is the container of accounts.
type Wallet struct {
	ID string
	s  Store
}

// init wallet
// 1. dail main chain, init chain interface and save to pool
// 2. init wallet store
func (w *Wallet)Init(chainUrl string, chainTypes []int) (error){
	for chainType := range(chainTypes){
		if chainType == int(chain.Eth){
			//dial eth chain
			url := chainUrl
			
			client, err := ethclient.Dial(url)
			if err != nil {
				return fmt.Errorf("failed to connect eth rpc endpoint {%v}, err is:%v \n", url, err)
			}
			blc = eth.NewChainEthereum(client)
			chain.Set(chain.Eth, blc)

		}
		if chainType == int(chain.EOS){
			fmt.Println("not support EOS yet.")
		}
	}

	w.s = &AccountRecord{}
	return nil
}

//UseAccount selects an account to user
func (w *Wallet) SetAccount(ctx context.Context, accountID string, pubKey string, chainType chain.ChainType) error {
	return w.s.SetAccountBasic(ctx, accountID, pubKey, chainType)
}


func (w *Wallet) IsExistAccount(ctx context.Context, accountID string) bool{
	return w.s.IsExistAccount(ctx, accountID)
}

//UseAccount selects an account to user
func (w *Wallet) UseAccount(ctx context.Context, accountID string) (*Account, error) {
	accountBasic, err := w.s.GetAccountBasic(ctx, accountID)
	buf, err := base64.StdEncoding.DecodeString(accountBasic.PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key, caused by:%v", err)
	}

	a := &Account{ID: accountBasic.ID, s: w.s}
	var blc chain.Chain
	if blc, err = chain.Get(chain.ChainType(accountBasic.ChainType)); err != nil {
		return nil, err
	}
	a.blc = blc

	// var pub crypto.PublicKey

	a.pubkey, err = gcrypto.UnmarshalPubkey(buf)
	if err != nil{
		return nil, fmt.Errorf("failed to parse public key, caused by:%v", err)
	}

	// if pub, err = x509.ParsePKIXPublicKey(buf); err != nil {
		// return nil, fmt.Errorf("failed to parse public key, caused by:%v", err)
	// }
	// if ecdsaPub, ok := pub.(*ecdsa.PublicKey); !ok {
		// return nil, fmt.Errorf("not a ecdsa public key, caused by:%v", err)
	// } else {
		// a.pubkey = ecdsaPub
	// }

	return a, nil
}

// Create Contract Transaction
func (w Wallet) CreateContentContractTrans(ctx context.Context, a *Account, data map[string]interface{}) (*transaction.ContractTransaction, error) {
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
func (w Wallet) AfterSign(ctx context.Context, a *Account, sigHex string, trans *transaction.ContractTransaction) (error) {
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
