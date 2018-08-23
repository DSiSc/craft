package types

import (
	"bytes"
	"github.com/stretchr/testify/assert"
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
