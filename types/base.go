/*
*  Define foundation of other types
 */

package types

import (
	"io"
)

// Lengths of hashes and addresses in bytes.
const (
	HashLength    = 32
	AddressLength = 20
)

// Address represents the 20 byte address of an Ethereum account.
type Address [AddressLength]byte

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

func (h *Hash) Serialize(w io.Writer) error {
	_, err := w.Write(h[:])
	return err
}
