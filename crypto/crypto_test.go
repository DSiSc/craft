package crypto

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testAddrHex = "970e8128ab834e8eac17ab8e3812f010678cf791"

func Test_CreateAddress(t *testing.T) {
	assert := assert.New(t)

	addr := types.HexToAddress(testAddrHex)
	caddr0 := CreateAddress(addr, 0)

	assert.Equal(types.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"), caddr0)
}
