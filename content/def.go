package content

import "github.com/icstglobal/go-icst/user"

type Content struct {
	Hash  []byte
	Data  []byte
	Owner *user.User
}
