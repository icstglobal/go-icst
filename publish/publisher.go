package publish

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"log"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/contract"
	"github.com/icstglobal/go-icst/skill"
	"github.com/icstglobal/go-icst/user"
)

// ContentPublisher is interface for user to publish a content
type ContentPublisher interface {
	Pub(c *content.Content)
}

// SkillPublisher is interface for user to publish a skill as a service
type SkillPublisher interface {
	Pub(s *skill.Skill)
}

//PubContent upload content hash to block chain and returns the address of smart contract
func PubContent(owner *user.User, c *content.Content, fee uint32, platform *user.User, ratio uint8) ([]byte, error) {
	hash := hash(c.Data)
	if bytes.Equal(hash, c.Hash) {
		return nil, errors.New("hash does not match")
		// return error.Error("hash doest not match")
	}

	ct := contract.NewConsumeContract(c, platform, owner, fee, ratio)
	//TODO:determine chain type
	chain, err := chain.NewChain(chain.Ethereum)
	ctAddr, err := chain.DeployContract(ct)
	if err != nil {
		log.Printf("faild to publish content:%v, error:%v\n", c, err)
	}

	//TODO:save metadata of content and smart contract
	return ctAddr, nil
}

func hash(data []byte) []byte {
	//TODO:change to another hash maybe
	sha := sha256.Sum256(data)
	return sha[:]
}
