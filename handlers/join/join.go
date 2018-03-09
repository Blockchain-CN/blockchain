// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package join

import (
	"github.com/Blockchain-CN/blockchain/common"
	idl "github.com/Blockchain-CN/blockchain/idls/join"
	pto "github.com/Blockchain-CN/blockchain/protocal"
)

// AddPeer join to the blockchain system by connect to a peer
func AddPeer(req *idl.JRequest) *idl.JResponse {
	resp := idl.NewJResponse()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	err := pto.AddPeer(req.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp
}
