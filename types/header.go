/*
*  Define types and structures relate header
 */
package types

import (
	"bytes"
	"crypto/sha256"
	"github.com/DSiSc/craft/serialize"
	"io"
)

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	PrevBlockHash Hash     // preblock hash
	StateRoot     Hash     // statedb root
	TxRoot        Hash     // transactions root
	ReceipsRoot   Hash     // receipt root
	Height        uint32   // block height
	Timestamp     uint32   // timestamp
	MixDigest     Hash     // digest
	SigData       [][]byte // signatures
	hash          *Hash    // header hash
}

//Serialize the blockheader data without program
func (h *Header) SerializeUnsigned(w io.Writer) error {
	err := h.PrevBlockHash.Serialize(w)
	if err != nil {
		return err
	}
	err = h.TxRoot.Serialize(w)
	if err != nil {
		return err
	}
	err = serialize.WriteUint32(w, h.Timestamp)
	if err != nil {
		return err
	}
	err = serialize.WriteUint32(w, h.Height)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) Hash() Hash {
	if h.hash != nil {
		return *h.hash
	}
	buf := new(bytes.Buffer)
	h.SerializeUnsigned(buf)
	temp := sha256.Sum256(buf.Bytes())
	hash := Hash(sha256.Sum256(temp[:]))

	h.hash = &hash
	return hash
}
