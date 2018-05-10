package content

import (
	"encoding/hex"
	"time"

	"github.com/icstglobal/go-icst/metadata"
)

type Content struct {
	Hash []byte
	Data []byte
	//address of the owner
	Owner []byte
}

//FromContent creates ContentMetadata instance from the content
func ToMetadata(c *Content) metadata.ContentMetadata {
	return metadata.ContentMetadata{
		Hash:  hex.EncodeToString(c.Hash),
		Date:  time.Now().Unix(),
		Owner: hex.EncodeToString(c.Owner),
	}
}
