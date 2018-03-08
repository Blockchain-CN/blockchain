// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package create

import (
	idl "github.com/Blockchain-CN/blockchain/idls/create"
)

// GenerateBlock create a new block and spread it.
func GenerateBlock(req *idl.CRequest) *idl.CResponse {
	resp := idl.NewResponseIDL()
	return resp
}
