package ethereum

import (
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/contract"
)

type ChainEthereum struct {
	Backend bind.ContractBackend
}

// GetContract gets smart contract from Ethereum chain with its address.
func (c *ChainEthereum) GetContract(addr []byte) (interface{}, error) {
	ct, err := NewConsumeContent(common.BytesToAddress(addr), c.Backend)
	if err != nil {
		log.Printf("faild to find smart contract, err:%v", err)
	}
	return ct, errors.New("not implemented yet")
}

// DeployContract method convert the domain contract to Ehtereum smart contract and deploy it.
// The address of the deployed smart contract will be returned if success.
func (c *ChainEthereum) DeployContract(contract *contract.ConsumeContract) (contractAddr []byte, err error) {
	opts := bind.NewKeyedTransactor(contract.Owner.PrivateKey)
	ownerAddr := ethcrypto.PubkeyToAddress(contract.Owner.PrivateKey.PublicKey)
	platformAddr := ethcrypto.PubkeyToAddress(contract.Platform.PrivateKey.PublicKey)
	if add, _, _, err := DeployConsumeContent(opts, c.Backend, ownerAddr, platformAddr, contract.Price, contract.Ratio); err != nil {
		// update contract address after deployed
		contract.Addr = add.Bytes()
		return nil, nil
	}

	return nil, err
}
