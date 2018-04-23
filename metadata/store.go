package metadata

//Store is a interface to save or retrieve metadat
type Store interface {
	Get(hash string) (interface{}, error)
	Save(hash string, data interface{}) error
}
