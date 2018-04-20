package metadata

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"github.com/icstglobal/go-icst/content"
	"github.com/icstglobal/go-icst/user"
)

func TestReadWriteRedis(t *testing.T) {
	var db Store
	//connect to local redis server
	db, err := NewDBRedis(nil)
	if err != nil {
		t.Fatal(err)
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal("failed to generate key,", err)
	}
	owner := &user.User{PrivateKey: key}
	ctt := &content.Content{Data: []byte("test data"), Owner: owner}
	hash := hashHex(ctt.Data)

	t.Log(hash)
	err = db.Save(hash, FromContent(ctt))
	if err != nil {
		t.Fatal("failed to save data,", err)
	}

	data, err := db.Get(hash)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(FromContent(ctt))
	if err != nil {
		t.Fatal(err)
	}
	metaExpected := string(buf)
	if strings.Compare(metaExpected, data.(string)) != 0 {
		t.Logf("expect %v, but get %v", metaExpected, data)
		t.FailNow()
	}
}

func hash(data []byte) []byte {
	//TODO:change to another hash maybe
	sha := sha256.Sum256(data)
	return sha[:]
}

func hashHex(data []byte) string {
	return hex.EncodeToString(data)
}
