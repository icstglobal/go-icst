package content

import "github.com/icstglobal/go-icst/user"

type Content struct {
	Hash  []byte
	Data  []byte
	owner *user.User
}

func (c *Content) BelongTo() user.User {
	return *c.owner
}
