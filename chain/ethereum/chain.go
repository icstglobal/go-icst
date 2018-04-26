package ethereum

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/contract"
)

type ChainEthereum struct {
	Backend *ethclient.Client
	// Client ethclient.Client
}

// GetContract gets smart contract from Ethereum chain with its address.
func (c *ChainEthereum) GetContract(addr []byte) (interface{}, error) {
	ct, err := NewConsumeContent(common.BytesToAddress(addr), c.Backend)
	if err != nil {
		log.Printf("faild to find smart contract, err:%v\n", err)
		return nil, err
	}
	return ct, nil
}

// DeployContract method convert the domain contract to Ehtereum smart contract and deploy it.
// The address of the deployed smart contract will be returned if success.
func (c *ChainEthereum) DeployContract(contract *contract.ConsumeContract) (contractAddr []byte, err error) {
	opts := bind.NewKeyedTransactor(contract.Owner.PrivateKey)
	ownerAddr := ethcrypto.PubkeyToAddress(contract.Owner.PrivateKey.PublicKey)
	platformAddr := ethcrypto.PubkeyToAddress(contract.Platform.PrivateKey.PublicKey)
	add, _, _, err := DeployConsumeContent(opts, c.Backend, ownerAddr, platformAddr, contract.Price, contract.Ratio)
	if err != nil {
		log.Printf("failed to deploy contract:%v\n", err)
		return nil, err
	}

	log.Printf("contract deployed to address:%+v\n", add.Hex())
	// update contract address after deployed
	contract.Addr = add.Bytes()
	return contract.Addr, err
}

//DeploySkillContract send the contract to block chain and wait for it to be mined.
//If the address returned is not nil, then it can be used even there is an error returned, but the contract may not yet be mined.
func (c *ChainEthereum) DeploySkillContract(ctx context.Context, contract *contract.SkillContract) (contractAddr []byte, err error) {
	opts := bind.NewKeyedTransactor(contract.Producuer.PrivateKey)
	prodAddr := ethcrypto.PubkeyToAddress(contract.Producuer.PrivateKey.PublicKey)
	platformAddr := ethcrypto.PubkeyToAddress(contract.Platform.PrivateKey.PublicKey)
	consAddr := ethcrypto.PubkeyToAddress(contract.Consumer.PrivateKey.PublicKey)
	hash := hex.EncodeToString(contract.Skill.Hash)
	addr, _, _, err := DeployConsumeSkill(opts, c.Backend, hash, prodAddr, platformAddr, consAddr, contract.Price, contract.Ratio)
	if err != nil {
		log.Printf("failed to deploy contract:%v\n", err)
		return nil, err
	}

	if err != nil {
		if err == bind.ErrNoCodeAfterDeploy {
			log.Println(err)
			return nil, err
		}

		log.Printf("waiting of deployment confirmation canceled:%v\n", err)
		//address of the contract can be used, as it has been sent to the chain
		return addr.Bytes(), err
	}

	// update contract address after deployed
	// contract.Addr = addr.Bytes()
	return contract.Addr, err
}
