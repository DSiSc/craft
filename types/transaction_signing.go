package types

import (
	"math/big"
)

type Signer interface {
	// Sender returns the sender address of the transaction.
	Sender(tx *Transaction) (Address, error)

	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error)

	// Hash returns the hash to be signed.
	Hash(tx *Transaction) Hash

	// Equal returns true if the given signer is the same as the receiver.
	Equal(Signer) bool
}

type sigCache struct {
	signer Signer
	from   Address
}

func Sender(signer Signer, tx *Transaction) (Address, error) {
	if sc := tx.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	addr, err := signer.Sender(tx)
	if err != nil {
		return Address{}, err
	}
	tx.from.Store(sigCache{signer: signer, from: addr})
	return addr, nil
}
