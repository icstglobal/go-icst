package transaction

type Transaction struct {
	rawTx               interface{}
	From                []byte
	To                  []byte
	TxHashFunc          func(rawTx interface{}) []byte
	TxHexHashSignedFunc func(rawTx interface{}) string
	SignFunc            func(sig []byte) error
}

func NewTransaction(rawTx interface{}, From []byte) *Transaction {
	ct := &Transaction{rawTx: rawTx, From: From}
	return ct
}

func (t *Transaction) Sender() []byte {
	cpy := make([]byte, len(t.From))
	copy(cpy, t.From)
	return cpy
}

func (t *Transaction) RawTx() interface{} {
	return t.rawTx
}
func (t *Transaction) SetRawTx(rawTx interface{}) {
	t.rawTx = rawTx
}

func (t *Transaction) Hash() []byte {
	return t.TxHashFunc(t.rawTx)
}

func (t *Transaction) Hex() string {
	return t.TxHexHashSignedFunc(t.rawTx)
}

func (t *Transaction) WithSign(sig []byte) error {
	return t.SignFunc(sig)
}
