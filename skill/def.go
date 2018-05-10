package skill

import (
	"github.com/icstglobal/go-icst/user"
)

// Skill is a kind of service to help others
type Skill struct {
	Hash     []byte
	Data     []byte
	Producer *user.User
}

type Contract struct {
	*Skill
	*Options
	Consumer *user.User
	Price    uint32
	Addr     []byte // the address of the smart contract after being deployed to the chain
	Nonce    uint64 // nonce of the transaction of the publisher's account
}

type Options struct {
	Platform *user.User
	Price    uint32
	Ratio    uint8
}

func NewContract(s *Skill, opts *Options, user *user.User) *Contract {
	return &Contract{s, opts, user, 0, nil, 0}
}
