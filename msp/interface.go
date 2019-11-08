package msp

import "github.com/DSiSc/craft/types"

// Msp member management interface
type Msp interface {
	// check whether address is authorized
	IsAuthorized(addr *types.Address) bool
}
