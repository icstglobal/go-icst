package consume

import (
	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/skill"
	"github.com/icstglobal/go-icst/user"
)

// ContentConsumer is the interface for user to consume a content
type ContentConsumer interface {
	Consume(u *user.User, c *content.Content)
}

// SkillConsumer is the interface for user to consume a skill service
type SkillConsumer interface {
	Consume(u *user.User, s *skill.Skill)
}
