package types

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_Hash(t *testing.T) {
	assert := assert.New(t)
	var txTemp Transaction
	tx := &txTemp
	hash := tx.Hash()
	assert.NotNil(hash)

	buf := bytes.Buffer{}
	err := tx.SerializeUnsigned(&buf)
	assert.Nil(err)
}

func Test_NewTransaction(t *testing.T) {
	assert := assert.New(t)
	b := Address{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}
	emptyTx := NewTransaction(
		0,
		b,
		big.NewInt(0), 0, big.NewInt(0),
		nil,
	)
	assert.NotNil(emptyTx)
}

type mockSign struct {
}

var b = Address{
	0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
	0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
}

func (*mockSign) Sender(tx *Transaction) (Address, error) {
	return b, nil
}

func (*mockSign) SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error) {
	return new(big.Int), new(big.Int), new(big.Int), nil
}

func (*mockSign) Hash(tx *Transaction) Hash {
	var temp Hash
	return temp
}

func (*mockSign) Equal(Signer) bool {
	return true
}

func NewMockSign() *mockSign {
	return &mockSign{}
}

func Test_AsMessage(t *testing.T) {
	assert := assert.New(t)
	b := Address{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}

	emptyTx := NewTransaction(
		0,
		b,
		big.NewInt(0), 0, big.NewInt(0),
		nil,
	)
	signer := NewMockSign()
	message, err := emptyTx.AsMessage(signer)
	assert.NotNil(message)
	assert.NotNil(message.from)
	assert.Nil(err)
}

func Test_NewMessage(t *testing.T) {
	assert := assert.New(t)
	//(from Address, to *Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, checkNonce bool)
	var data []byte = []byte("solo_node")
	message := NewMessage(b, &b, 0, big.NewInt(0), 0, big.NewInt(0), data, true)
	assert.NotNil(message)
	assert.Equal(b, message.From())
	assert.Equal(&b, message.To())
	assert.Equal(big.NewInt(0), message.GasPrice())
	assert.Equal(big.NewInt(0), message.Value())
	assert.Equal(uint64(0), message.Gas())
	assert.Equal(data, message.Data())
	assert.Equal(true, message.CheckNonce())
}

func Test_GetTxsRoot(t *testing.T) {
	assert := assert.New(t)
	var txs []*Transaction
	emptyTx := NewTransaction(
		0,
		b,
		big.NewInt(0), 0, big.NewInt(0),
		nil,
	)
	txs = append(txs, emptyTx)
	hash := GetTxsRoot(txs)
	var expect = Hash{
		0x81, 0xfc, 0x49, 0x25, 0x61, 0xda, 0x56, 0x83, 0x2f, 0x9a, 0x3c, 0xe1, 0xd0, 0x56, 0x9e, 0xa1,
		0x0, 0xc7, 0x69, 0x49, 0x54, 0x5d, 0x3b, 0x62, 0x54, 0xf1, 0x8, 0x6a, 0x33, 0xb7, 0x14, 0x13,
	}
	assert.Equal(expect, hash)
}
