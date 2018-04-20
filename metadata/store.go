package metadata

type Store interface {
	Get(hash string) (interface{}, error)
	Save(hash string, data interface{}) error
}
