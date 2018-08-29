/*
*  Define types and structures relate header
 */
package types

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"github.com/DSiSc/craft/serialize"
	"io"
)

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ChainID       uint64   `json:"chainId"    gencodec:"required"`     // chainid
	PrevBlockHash Hash     `json:"prevHash"    gencodec:"required"`    // preblock hash
	StateRoot     Hash     `json:"stateRoot"    gencodec:"required"`   // statedb root
	TxRoot        Hash     `json:"txRoot"    gencodec:"required"`      // transactions root
	ReceiptsRoot  Hash     `json:"receipsRoot"    gencodec:"required"` // receipt root
	Height        uint64   `json:"height"    gencodec:"required"`      // block height
	Timestamp     uint64   `json:"timestamp"    gencodec:"required"`   // timestamp
	MixDigest     Hash     `json:"mixDigest"    gencodec:"required"`   // digest
	SigData       [][]byte `json:"sigData"    gencodec:"required"`     // signatures
	BlockHash     *Hash    `json:"blockHash"    gencodec:"required"`   // block hash
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
	err = serialize.WriteUint64(w, h.Timestamp)
	if err != nil {
		return err
	}
	err = serialize.WriteUint64(w, h.Height)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) Hash() Hash {
	if h.BlockHash != nil {
		return *h.BlockHash
	}
	buf := new(bytes.Buffer)
	h.SerializeUnsigned(buf)
	temp := sha256.Sum256(buf.Bytes())
	hash := Hash(sha256.Sum256(temp[:]))

	h.BlockHash = &hash
	return hash
}

//Serialize the blockheader
func (h *Header) Serialize(w io.Writer) error {
	h.SerializeUnsigned(w)
	err := serialize.WriteVarUint(w, uint64(len(h.SigData)))
	if err != nil {
		return errors.New("serialize sig pubkey length failed")
	}

	for _, sig := range h.SigData {
		err = serialize.WriteVarBytes(w, sig)
		if err != nil {
			return err
		}
	}

	return nil
}
