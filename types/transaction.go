package types

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"github.com/DSiSc/craft/serialize"
	"io"
	"math/big"
	"sync/atomic"
)

type Transaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdata struct {
	AccountNonce uint64   `json:"nonce"    gencodec:"required"`
	Price        *big.Int `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64   `json:"gas"      gencodec:"required"`
	Recipient    *Address `json:"to"       rlp:"nil"`
	Amount       *big.Int `json:"value"    gencodec:"required"`
	Payload      []byte   `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *Hash `json:"hash" rlp:"-"`
}

func (t Transaction) AccountNonce() uint64 { return t.data.AccountNonce }
func (t Transaction) Price() *big.Int      { return t.data.Price }
func (t Transaction) GasLimit() uint64     { return t.data.GasLimit }
func (t Transaction) Recipient() *Address  { return t.data.Recipient }
func (t Transaction) Amount() *big.Int     { return t.data.Amount }
func (t Transaction) Payload() []byte      { return t.data.Payload }

func (tx *Transaction) Hash() Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(Hash)
	}
	buf := bytes.Buffer{}
	tx.SerializeUnsigned(&buf)
	temp := sha256.Sum256(buf.Bytes())
	f := Hash(sha256.Sum256(temp[:]))
	tx.hash.Store(f)
	return f
}

//Serialize the Transaction data without contracts
func (tx *Transaction) SerializeUnsigned(w io.Writer) error {
	//txType
	if err := serialize.WriteUint64(w, tx.data.AccountNonce); err != nil {
		return errors.New("[SerializeUnsigned], Transaction nonce failed.")
	}
	if err := serialize.WriteUint64(w, tx.data.GasLimit); err != nil {
		return errors.New("[SerializeUnsigned], Transaction nonce failed.")
	}

	return nil
}

func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

// New a transaction
func newTransaction(nonce uint64, to *Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	if len(data) > 0 {
		data = CopyBytes(data)
	}
	d := txdata{
		AccountNonce: nonce,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		GasLimit:     gasLimit,
		Price:        new(big.Int),
		V:            new(big.Int),
		R:            new(big.Int),
		S:            new(big.Int),
	}
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}

	return &Transaction{data: d}
}

func NewTransaction(nonce uint64, to Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	return newTransaction(nonce, &to, amount, gasLimit, gasPrice, data)
}

// Message evm context message
type Message struct {
	to         *Address
	from       Address
	nonce      uint64
	amount     *big.Int
	gasLimit   uint64
	gasPrice   *big.Int
	data       []byte
	checkNonce bool
}

func NewMessage(from Address, to *Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, checkNonce bool) Message {
	return Message{
		from:       from,
		to:         to,
		nonce:      nonce,
		amount:     amount,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
		data:       data,
		checkNonce: checkNonce,
	}
}

func (m Message) From() Address { return m.from }
func (m Message) To() *Address  { return m.to }
func (m Message) GasPrice() *big.Int  { return m.gasPrice }
func (m Message) Value() *big.Int     { return m.amount }
func (m Message) Gas() uint64         { return m.gasLimit }
func (m Message) Nonce() uint64       { return m.nonce }
func (m Message) Data() []byte        { return m.data }
func (m Message) CheckNonce() bool    { return m.checkNonce }