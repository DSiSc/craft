/*
*  Define types and structures relate header
 */
package types

import "math/big"

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ChainID       uint64  `json:"chainId"    gencodec:"required"`     // chainid
	PrevBlockHash Hash    `json:"prevHash"    gencodec:"required"`    // preblock hash
	StateRoot     Hash    `json:"stateRoot"    gencodec:"required"`   // statedb root
	TxRoot        Hash    `json:"txRoot"    gencodec:"required"`      // transactions root
	ReceiptsRoot  Hash    `json:"receiptsRoot"    gencodec:"required"` // receipt root
	Height        uint64  `json:"height"    gencodec:"required"`      // block height
	Timestamp     uint64  `json:"timestamp"    gencodec:"required"`   // timestamp
	CoinBase      Address `json:"coinbase"    gencodec:"required"`    // coin base

	// not contain when compute header hash
	MixDigest Hash     `json:"mixDigest"    gencodec:"required"` // digest
	SigData   [][]byte `json:"signData"    gencodec:"required"`  // SigData

	//match etherum
	ParentHash  Hash    `json:"parentHash"       gencodec:"required"`
	Coinbase    Address `json:"miner"            gencodec:"required"`
	TxHash      Hash    `json:"transactionsRoot" gencodec:"required"`
	UncleHash   Hash    `json:"sha3Uncles"    gencodec:"required"`
	Bloom       Bloom   `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	Nonce       uint64         `json:"nonce"            gencodec:"required"`
}
