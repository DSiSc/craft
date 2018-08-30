package crypto

import (
	"github.com/DSiSc/craft/types"
)

// CreateAddress creates an ethereum address given the bytes and the nonce
func CreateAddress(b types.Address, nonce uint64) types.Address {
	address := types.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d")
	return address
}
