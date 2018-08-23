/*
*  Define types and structures relate header
 */
package types

import (
	"math/big"
)

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	PreHash     Hash     // preblock hash
	StateRoot   Hash     // statedb root
	TxRoot      Hash     // transactions root
	ReceipsRoot Hash     // receipt root
	Height      *big.Int // block height
	Timestamp   *big.Int // timestamp
	MixDigest   Hash     // digest
	SigData     [][]byte // signatures
	Hash        Hash     // header hash
}
