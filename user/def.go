package user

import (
	"crypto/ecdsa"
)

// User is the people of real world. They can publish, consume or judge a content or skill.
type User struct {
	PrivateKey *ecdsa.PrivateKey
}
