package publish

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"log"

	"github.com/icstglobal/go-icst/metadata"
	"github.com/icstglobal/go-icst/skill"

	"github.com/icstglobal/go-icst/chain"
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/contract"
	"github.com/icstglobal/go-icst/user"
)

// ContentPublisher is interface for user to publish a content
type ContentPublisher struct {
	chain chain.Chain
}

func NewContentPublisher(chain chain.Chain) *ContentPublisher {
	return &ContentPublisher{chain: chain}
}

//PubContent upload content hash to block chain and returns the address of smart contract
func (p *ContentPublisher) PubContent(c *content.Content, opts contract.Options) ([]byte, error) {
	hash := hash(c.Data)
	if bytes.Equal(hash, c.Hash) {
		return nil, errors.New("hash does not match")
	}

	ct := contract.NewConsumeContract(c, opts.Platform, opts.Price, opts.Ratio)
	ctAddr, err := p.chain.DeployContract(ct)
	if err != nil {
		log.Printf("faild to publish content:%v, error:%v\n", c, err)
		return nil, err
	}

	//TODO:save metadata of content and smart contract
	return ctAddr, nil
}

func hash(data []byte) []byte {
	//TODO:change to another hash maybe
	sha := sha256.Sum256(data)
	return sha[:]
}

// SkillPublisher is interface for user to publish a skill as a service
type SkillPublisher struct {
	chain.Chain
	store metadata.Store
}

func NewSkillPublisher(chain chain.Chain, store metadata.Store) *SkillPublisher {
	return &SkillPublisher{chain, store}
}

func (p *SkillPublisher) Pub(ctx context.Context, s *skill.Skill, opts *contract.Options, consumer *user.User) (addr []byte, err error) {
	ct := contract.NewSkillContract(s, opts, consumer)
	addr, err = p.Chain.DeployContract(ctx, ct)
	if err != nil {
		log.Println("failed to publish skill contract, ", err)
		return nil, err
	}
	//TODO: save metadata
	return
}
