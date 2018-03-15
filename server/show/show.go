// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package show

import (
	handler "github.com/Blockchain-CN/blockchain/handlers/show"
	idl "github.com/Blockchain-CN/blockchain/idls/show"
)

// JController ...
type SController struct {
}

// GenIdl ...
func (c *SController) GenIdl() interface{} {
	return idl.NewJRequest()
}

// Do ...
func (c *SController) Do(req interface{}) interface{} {
	r := req.(*idl.SRequest)
	return handler.Show(r)
}
