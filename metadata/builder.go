package metadata

import (
	"encoding/json"
)

//ContentMetadata represents metadata of a content
type ContentMetadata struct {
	Hash  string `json:"hash,required"`
	Date  int64  `json:"date,required"`
	Owner string `json:"owner,required"`
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

//SkillMetadata represents metadata of a skill contract
type SkillMetadata struct {
	Hash      string `json:"hash,required"`
	Date      int64  `json:"date,required"`
	Price     uint32 `json:"price,required"`
	Publisher string `json:"publisher,required"`
	Platform  string `json:"platform,required"`
	Consumer  string `json:"consumer,required"`
}
