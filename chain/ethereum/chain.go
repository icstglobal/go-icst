package ethereum

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icstglobal/go-icst/contract"
)

//ContractDeploymentTimeout  is the default timeout for contract deployment
const ContractDeploymentTimeout = 20 * time.Second

// ChainEthereum is a wrapper for ethereum client
type ChainEthereum struct {
	contractBackend bind.ContractBackend
	deployBackend   bind.DeployBackend
}

// NewChainEthereum creates a new Ethereum chain object
func NewChainEthereum(client *ethclient.Client) *ChainEthereum {
	return &ChainEthereum{
		contractBackend: client,
		deployBackend:   client,
	}
}

// GetContentContract gets content contract from Ethereum chain with its address.
func (c *ChainEthereum) getContentContract(addr []byte) (*ConsumeContent, error) {
	ct, err := NewConsumeContent(common.BytesToAddress(addr), c.contractBackend)
	if err != nil {
		log.Printf("faild to find smart contract, err:%v\n", err)
		return nil, err
	}
	return ct, nil
}

// GetSkillContract gets skill contract from Ethereum chain with its address.
func (c *ChainEthereum) getSkillContract(addr []byte) (*ConsumeSkill, error) {
	ct, err := NewConsumeSkill(common.BytesToAddress(addr), c.contractBackend)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

// GetContract gets smart contract from Ethereum chain with its address.
func (c *ChainEthereum) GetContract(addr []byte, t reflect.Type) (interface{}, error) {
	switch t {
	case reflect.TypeOf((*contract.ConsumeContract)(nil)):
		return c.getContentContract(addr)
	case reflect.TypeOf((*contract.SkillContract)(nil)):
		return c.getSkillContract(addr)
	default:
		return nil, fmt.Errorf("unknown contract type:%v", t.Name())
	}
}

// DeployContract method convert the domain contract to Ehtereum smart contract and deploy it.
// The address of the deployed smart contract will be returned if success.
func (c *ChainEthereum) DeployContract(ctx context.Context, icontract interface{}) (contractAddr []byte, err error) {
	switch icontract.(type) {
	case *contract.ConsumeContract:
		return c.deployContentContract(context.TODO(), icontract.(*contract.ConsumeContract))
	case *contract.SkillContract:
		return c.deploySkillContract(ctx, icontract.(*contract.SkillContract))
	default:
		return nil, errors.New("unsupported contract type")
	}
}

func (c *ChainEthereum) deployContentContract(ctx context.Context, contract *contract.ConsumeContract) (contractAddr []byte, err error) {
	opts := bind.NewKeyedTransactor(contract.Owner.PrivateKey)
	ownerAddr := ethcrypto.PubkeyToAddress(contract.Owner.PrivateKey.PublicKey)
	platformAddr := ethcrypto.PubkeyToAddress(contract.Platform.PrivateKey.PublicKey)
	addr, tx, _, err := DeployConsumeContent(opts, c.contractBackend, ownerAddr, platformAddr, contract.Price, contract.Ratio)
	if err != nil {
		return nil, err
	}

	if addr, err = c.waitContractDeployed(ctx, tx); err == bind.ErrNoCodeAfterDeploy {
		return nil, err
	}
	// update contract address after deployed
	contract.Addr = addr.Bytes()
	contract.Nonce = tx.Nonce()
	return contract.Addr, err
}

//DeploySkillContract send the contract to block chain and wait for it to be mined.
//If the address returned is not nil, then it can be used even there is an error returned, but the contract may not yet be mined.
func (c *ChainEthereum) deploySkillContract(ctx context.Context, contract *contract.SkillContract) (contractAddr []byte, err error) {
	opts := bind.NewKeyedTransactor(contract.Producer.PrivateKey)
	prodAddr := ethcrypto.PubkeyToAddress(contract.Producer.PrivateKey.PublicKey)
	platformAddr := ethcrypto.PubkeyToAddress(contract.Platform.PrivateKey.PublicKey)
	consAddr := ethcrypto.PubkeyToAddress(contract.Consumer.PrivateKey.PublicKey)
	hash := hex.EncodeToString(contract.Skill.Hash)
	addr, tx, _, err := DeployConsumeSkill(opts, c.contractBackend, hash, prodAddr, platformAddr, consAddr, contract.Price, contract.Ratio)
	if err != nil {
		return nil, err
	}

	if addr, err = c.waitContractDeployed(ctx, tx); err == bind.ErrNoCodeAfterDeploy {
		return nil, err
	}
	// update contract address after deployed
	contract.Addr = addr.Bytes()
	contract.Nonce = tx.Nonce()
	return contract.Addr, err
}

func (c *ChainEthereum) waitContractDeployed(ctx context.Context, tx *types.Transaction) (common.Address, error) {
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), ContractDeploymentTimeout)
	}

	return bind.WaitDeployed(ctx, c.deployBackend, tx)
}

func (c *ChainEthereum) WatchEvent(ctx context.Context, contractDeployed *ConsumeSkill, stateChan chan<- *ConsumeSkillStateChange) (event.Subscription, error) {
	watchOpts := &bind.WatchOpts{Start: nil, Context: ctx} // start from the latest block
	return contractDeployed.WatchStateChange(watchOpts, stateChan)
}
