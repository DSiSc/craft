package msp

import "github.com/DSiSc/craft/types"

type PermitAllMsp struct {
}

func NewPermitAllMsp() *PermitAllMsp {
	return &PermitAllMsp{}
}

func (self *PermitAllMsp) IsAuthorized(addr *types.Address) bool {
	return true
}
