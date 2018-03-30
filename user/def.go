package user

// User is the people of real world. They can publish, consume or judge a content or skill.
type User struct {
	PrivateKey []byte
	PublicKey  []byte
}

// Addr returns the transaction address of the user
func (u *User) Addr() []byte {
	// TODO: can generate from private key
	return u.PublicKey
}
