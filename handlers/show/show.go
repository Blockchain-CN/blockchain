// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package show

import (
	"github.com/Blockchain-CN/blockchain/models"
	idl "github.com/Blockchain-CN/blockchain/idls/show"
	pto "github.com/Blockchain-CN/blockchain/protocal"
)

type result struct {

}

// Show join to the blockchain system by connect to a peer
func Show(req *idl.SRequest) *idl.SResponse {
	resp := idl.NewJResponse()
	single := pto.GetProtocal()
	if req.Chain {
		resp.Chain = models.FetchChain()
	}
	if req.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp
}
