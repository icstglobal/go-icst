package common

import (
	"crypto/sha256"
)

//Hash returns the hash of input
func Hash(input []byte) [32]byte {
	return sha256.Sum256(input)
}
