package types

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ComputeMerkleRoot(t *testing.T) {
	var data []Hash
	a1 := Hash(sha256.Sum256([]byte("a")))
	a2 := Hash(sha256.Sum256([]byte("b")))
	a3 := Hash(sha256.Sum256([]byte("c")))
	a4 := Hash(sha256.Sum256([]byte("d")))
	a5 := Hash(sha256.Sum256([]byte("e")))
	data = append(data, a1)
	data = append(data, a2)
	data = append(data, a3)
	data = append(data, a4)
	data = append(data, a5)
	hash := ComputeMerkleRoot(data)
	assert.NotEqual(t, hash, Hash{})
}
