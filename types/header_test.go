package types

import (
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
}
