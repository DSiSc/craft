package types

import (
	"math/big"
	"sync/atomic"
)

type Transaction struct {
	Data TxData
	Hash atomic.Value
	Size atomic.Value
	From atomic.Value
}

type TxData struct {
	AccountNonce uint64   `json:"nonce"    gencodec:"required"`
	Price        *big.Int `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64   `json:"gas"      gencodec:"required"`
	Recipient    *Address `json:"to"       rlp:"nil"`
	From         *Address `json:"from"     rlp:"nil"`
	Amount       *big.Int `json:"value"    gencodec:"required"`
	Payload      []byte   `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *Hash `json:"hash" rlp:"-"`
}
