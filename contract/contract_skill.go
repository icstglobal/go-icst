package contract

import (
	"github.com/icstglobal/go-icst/skill"
	"github.com/icstglobal/go-icst/user"
)

type SkillContract struct {
	*skill.Skill
	*Options
	Consumer *user.User
	Price    uint32
	Addr     []byte // the address of the smart contract after being deployed to the chain
	Nonce    uint64 // nonce of the transaction of the publisher's account
}

func NewSkillContract(s *skill.Skill, opts *Options, user *user.User) *SkillContract {
	return &SkillContract{s, opts, user, 0, nil, 0}
}
