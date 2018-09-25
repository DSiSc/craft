/*
*  Define types and structures relate block
 */
package types

type Block struct {
	Header       *Header
	Transactions []*Transaction
	HeaderHash   Hash `json:"headerHash"    gencodec:"required"`
}
