package types

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader_Serialize(t *testing.T) {
	assert := assert.New(t)
	header := Header{}
	header.Height = 321
	header.SigData = make([][]byte, 0)
	hash := header.Hash()
	assert.NotNil(hash)

	buf := bytes.NewBuffer(nil)
	err := header.Serialize(buf)
	bs := buf.Bytes()
	assert.Nil(err)
	assert.NotNil(bs)
}
