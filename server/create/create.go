// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package create

import (
	handler "github.com/Blockchain-CN/blockchain/handlers/create"
	idl "github.com/Blockchain-CN/blockchain/idls/create"
)

// CController ...
type CController struct {
}

// GenIdl ...
func (c *CController) GenIdl() interface{} {
	return idl.NewRequestIDL()
}

// Do ...
func (c *CController) Do(req interface{}) interface{} {
	r := req.(*idl.CRequest)
	return handler.GenerateBlock(r)
}
