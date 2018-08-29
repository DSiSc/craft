package types

import (
	"github.com/DSiSc/craft/signature/keypair"
)

type Account struct {
	PrivateKey keypair.PrivateKey
	PublicKey  keypair.PublicKey
	Address    Address
	SigScheme  keypair.SignatureScheme
}
