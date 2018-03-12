// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package create

import (
	"github.com/Blockchain-CN/blockchain/common"
	idl "github.com/Blockchain-CN/blockchain/idls/create"
	pto "github.com/Blockchain-CN/blockchain/protocal"
)

// GenerateBlock create a new block and spread it.
func GenerateBlock(req *idl.CRequest) *idl.CResponse {
	resp := idl.NewCResponseIDL()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	pto.DataQueue <- req
	return resp
}
