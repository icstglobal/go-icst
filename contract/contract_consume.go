package contract

import (
	"errors"

	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/user"
)

// ConsumeContract is the contract for consuming a content directly
type ConsumeContract struct {
	Ctt      *content.Content
	Platform []byte
	Owner    []byte
	Price    uint32
	Ratio    uint8  // ratio to split the fee
	Addr     []byte // the address of the smart contract after being deployed to the chain
	Nonce    uint64
}

// NewConsumeContract creates a new contract for content consuming. It can never be changed.
func NewConsumeContract(c *content.Content, platform []byte,
	fee uint32, ratio uint8) *ConsumeContract {
	return &ConsumeContract{Ctt: c, Platform: platform, Owner: c.Owner,
		Price: fee, Ratio: ratio,
	}
}

//Exec a contract and the user pay the fee
func (cc *ConsumeContract) Exec(u *user.User) error {
	// a user consume a content, so the user needs to send fee from the user's address
	// to the smart contract's address
	return nil
}

// Deploy the contract to underlying chain.
func (cc *ConsumeContract) Deploy() error {
	//TODO:do more contract validation
	if cc.Ctt == nil || cc.Platform == nil || cc.Owner == nil {
		return errors.New("content, platform or owner can not be empty")
	}

	if cc.Price > 0 && cc.Ratio == 0 {
		return errors.New("ratio should be set when fee is not zero")
	}
	//TODO:call chain API
	return nil
}

// Destory a contract will makes the contract dead on the chain
func (cc *ConsumeContract) Destory(holders []*user.User) error {
	// owners of the contract decide together to close the contract
	// TODO:validate the holders of the contract first
	//TODO:call chain API
	return nil
}
