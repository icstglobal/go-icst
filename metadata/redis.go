package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//DBRedis is a Redis storage implementation of store interface
type DBRedis struct {
	conn redis.Conn
}

//NewDBRedis create a new DBRedis instance.
//Pass nil as url to connect to local redis server running on default.
func NewDBRedis(url *string) (*DBRedis, error) {
	var conn redis.Conn
	var err error
	if url == nil {
		conn, err = redis.Dial("tcp", ":6379")

	} else {
		conn, err = redis.DialURL(*url)
	}

	if err != nil {
		return nil, err
	}

	return &DBRedis{conn: conn}, nil
}

//Get metadata by hash key
func (db *DBRedis) Get(hash string) (v interface{}, err error) {
	//currently we store all data as string into Redis
	return redis.String(db.conn.Do("GET", hash))
}

//Save metadata to the hash key
func (db *DBRedis) Save(hash string, data interface{}) (err error) {
	var v string
	switch data.(type) {
	case string:
		v = data.(string)
		break
	case []byte:
		v = string(data.([]byte))
		break
	default:
		//marshal to json
		var buf []byte
		if buf, err = json.Marshal(data); err != nil {
			return fmt.Errorf("data can not be serialized as json, inner error is:%v", err)
		}
		v = string(buf)
		break
	}
	//always overwrite the existing data
	_, err = db.conn.Do("SET", hash, v)
	return err
}
