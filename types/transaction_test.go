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
