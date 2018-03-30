package publish

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/skill"
	"github.com/icstglobal/go-icst/transaction"
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

func PubContent(owner *user.User, c *content.Content) error {
	hash := hash(c.Data)
	if bytes.Equal(hash, c.Hash) {
		return errors.New("hash does not match")
		// return error.Error("hash doest not match")
	}

	//TODO:generate transaction and save to underlying chain
	tx := transaction.NewContent(owner, c.Data)
	//TODO: sign the trans
	return nil
}

func hash(data []byte) []byte {
	//TODO:change to another hash maybe
	sha := sha256.Sum256(data)
	return sha[:]
}
