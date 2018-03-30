package transaction

import (
	"github.com/icstglobal/go-icst/user"
)

// NewContent create a new transaction to record a content
func NewContent(u *user.User, data []byte) Transaction {
	var trans Transaction
	trans.Data = data
	trans.From = u.Addr()
	return trans
}
