package content

import "github.com/icstglobal/go-icst/user"

type Content struct {
	Hash     []byte
	Data     []byte
	belongTo *user.User
}

func (c *Content) BelongTo() user.User {
	return *c.belongTo
}
