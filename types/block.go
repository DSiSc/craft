/*
*  Define types and structures relate block
 */
package types

type Block struct {
	Header       *Header
	Transactions []*Transaction
	HeaderHash   Hash     `json:"headerHash"    gencodec:"required"`
	SigData      [][]byte `json:"signData"    gencodec:"required"`
	BlockHash    Hash     `json:"blockHash"    gencodec:"required"`
}
