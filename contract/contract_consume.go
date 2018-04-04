package contract

import (
	"errors"

	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/user"
)

// ConsumeContract is the contract for consuming a content directly
type ConsumeContract struct {
	ctt       *content.Content
	platform  *user.User
	publisher *user.User
	fee       uint32
	ratio     uint8  // ratio to split the fee
	addr      []byte // the address of the smart contract after being deployed to the chain
}

// NewConsumeContract creates a new contract for content consuming. It can never be changed.
func NewConsumeContract(c *content.Content, platform *user.User, publisher *user.User,
	fee uint32, ratio uint8) *ConsumeContract {
	return &ConsumeContract{ctt: c, platform: platform, publisher: publisher,
		fee: fee, ratio: ratio,
	}
}

// Addr returns a copy of the contract's address.
// nil if the contract has never been deployed.
func (cc *ConsumeContract) Addr() []byte {
	if cc.addr == nil {
		return nil
	}
	// make a copy so that the address won't be modified by accident
	addrCopy := make([]byte, len(cc.addr))
	copy(cc.addr, addrCopy)
	return addrCopy
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
	if cc.ctt == nil || cc.platform == nil || cc.publisher == nil {
		return errors.New("content, platform or publisher can not be empty")
	}

	if cc.fee > 0 && cc.ratio == 0 {
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
