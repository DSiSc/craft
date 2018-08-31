/*
*  Define foundation of other types
 */

package types

// Lengths of hashes and addresses in bytes.
const (
	HashLength    = 32
	AddressLength = 20
)

const (
	// BloomByteLength represents the number of bytes used in a header log bloom.
	BloomByteLength = 256
)

// StorageSize is a wrapper around a float value that supports user friendly
// formatting.
type StorageSize float64

// Type to mark uniqueness of a node
type NodeAddress string

// Address represents the 20 byte address of an Ethereum account.
type Address [AddressLength]byte

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomByteLength]byte
