package transaction

type ContractTransaction struct {
	rawTx        interface{}
	sender       []byte
	ContractAddr []byte
	TxHashFunc   func(rawTx interface{}) []byte
	TxHexHashSignedFunc   func(rawTx interface{}) string
	SignFunc     func(sig []byte) error
}

func NewContractTransaction(rawTx interface{}, sender []byte) *ContractTransaction {
	ct := &ContractTransaction{rawTx: rawTx, sender: sender}
	return ct
}

func (t *ContractTransaction) Sender() []byte {
	cpy := make([]byte, len(t.sender))
	copy(cpy, t.sender)
	return cpy
}

func (t *ContractTransaction) RawTx() interface{} {
	return t.rawTx
}
func (t *ContractTransaction) SetRawTx(rawTx interface{}) {
	t.rawTx = rawTx
}

func (t *ContractTransaction) Hash() []byte {
	return t.TxHashFunc(t.rawTx)
}

func (t *ContractTransaction) Hex() string {
	return t.TxHexHashSignedFunc(t.rawTx)
}

func (t *ContractTransaction) WithSign(sig []byte) error {
	return t.SignFunc(sig)
}
