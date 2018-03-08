// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package join

import (
	handler "github.com/Blockchain-CN/blockchain/handlers/join"
	idl "github.com/Blockchain-CN/blockchain/idls/join"
)

// JController ...
type JController struct {
}

// GenIdl ...
func (c *JController) GenIdl() interface{} {
	return idl.NewJRequest()
}

// Do ...
func (c *JController) Do(req interface{}) interface{} {
	r := req.(*idl.JRequest)
	return handler.AddPeer(r)
}
