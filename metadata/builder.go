package metadata

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/icstglobal/go-icst/content"
)

//ContentMetadata represents metadata of a content
type ContentMetadata struct {
	Hash  string `json:"hash,required"`
	Date  int64  `json:"date,required"`
	Owner UserID `json:"owner,required"`
}

// UserID identify a user by its public key
type UserID struct {
	PublicKey crypto.PublicKey `json:"public_key,required"`
	//crypto algorithm to apply the publick key, like AES, ecdsa
	Crypto string `json:"crypto,required"`
}

//FromContent creates ContentMetadata instance from the content
func FromContent(c *content.Content) ContentMetadata {
	return ContentMetadata{
		Hash:  hex.EncodeToString(c.Hash),
		Date:  time.Now().Unix(),
		Owner: UserID{c.Owner.PrivateKey.Public(), "ecdsa"},
	}
}

//ContentMetaFromJSON deserialize content metadata from json string
func ContentMetaFromJSON(metaString string) (meta ContentMetadata, err error) {
	err = json.Unmarshal([]byte(metaString), meta)
	return
}

//ContentMetaToJSON serialize content metadata to json string
func ContentMetaToJSON(meta ContentMetadata) (string, error) {
	var buf []byte
	var err error
	if buf, err = json.Marshal(meta); err != nil {
		return "", err
	}

	return string(buf), nil
}

type SkillMetadata struct {
	Hash      string `json:"hash,required"`
	Date      int64  `json:"date,required"`
	Price     uint32 `json:"price,required"`
	Publisher UserID `json:"publisher,required"`
	Platform  UserID `json:"platform,required"`
	Consumer  UserID `json:"consumer,required"`
}
