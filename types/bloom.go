package types

const (
	// BloomByteLength represents the number of bytes used in a header log bloom.
	BloomByteLength = 256
)

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomByteLength]byte
