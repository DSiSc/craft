/*
*  Define types and structures relate block
 */
package types

type Block struct {
	Header       *Header
	Transactions []*Transaction
}
