package transaction

type Transaction struct {
	From []byte
	To   []byte
	Data []byte
	Fee  int32 // token or money costed
}
