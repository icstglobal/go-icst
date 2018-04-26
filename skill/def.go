package skill

import "github.com/icstglobal/go-icst/user"

// Skill is a kind of service to help others
type Skill struct {
	Hash      []byte
	Data      []byte
	Producuer *user.User
}
