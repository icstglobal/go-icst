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
)

// ContentPublisher is interface for user to publish a content
type ContentPublisher struct {
	chain chain.Chain
}

// SkillPublisher is interface for user to publish a skill as a service
type SkillPublisher interface {
	Pub(s *skill.Skill)
}

func NewContentPublisher(chain chain.Chain) *ContentPublisher {
	return &ContentPublisher{chain: chain}
}

//PubContent upload content hash to block chain and returns the address of smart contract
func (p *ContentPublisher) PubContent(c *content.Content, opts contract.Options) ([]byte, error) {
	hash := hash(c.Data)
	if bytes.Equal(hash, c.Hash) {
		return nil, errors.New("hash does not match")
		// return error.Error("hash doest not match")
	}

	ct := contract.NewConsumeContract(c, opts.Platform, opts.Price, opts.Ratio)
	ctAddr, err := p.chain.DeployContract(ct)
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
