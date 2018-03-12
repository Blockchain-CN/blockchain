// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package protocal

import (
	"time"
	"encoding/json"
	"sync"
	"fmt"

	p2p "github.com/Blockchain-CN/pheromones"
	"github.com/Blockchain-CN/blockchain/models"
	idl "github.com/Blockchain-CN/blockchain/idls/create"
)

var (
	singleton *Protocal
	DataQueue chan idl.CRequest
	wg sync.WaitGroup
	ip string
)

// InitPto init the default protocal object
func InitPto(addr string, to time.Duration) {
	r1 := p2p.NewSRouter(to)
	p1 := NewProtocal(addr, r1, to)
	s1 := p2p.NewServer(p1, to)
	singleton = p1
	DataQueue = make(chan string, 100)
	println("P2P Servering on ", addr)
	go BlockPublisher()
	go s1.ListenAndServe(addr)
}

// AddPeer add a peer to the default protocal's router
func AddPeer(addr string) error {
	return singleton.Add(addr, addr)
}

//
func BlockPublisher() {
	for {
		select {
		case ud := <- DataQueue:
			// get user object
			user, err := models.Login(ud.Name)
			if err != nil {
				return err
			}

			// get trans object
			trans, err := models.GenerateTransWithKey(user.Public, user.Private, ud.Data)
			if err != nil {
				return err
			}
			transStr, err := json.Marshal(trans)
			if err != nil {
				return err
			}

			// append a block to the chain until succeed
			for {
				// get block object
				block := models.GenerateBlock(models.GetChainTail().Hash, string(transStr), models.GetChainLen())

				// add to blockchain
				err = models.AppendChain(block)
				if err == nil {
					spreads(block)
					break
				}
			}
		}
	}
}

// spread the latest block to all peers
func spreads(block *models.Block) {
	blockStr, err := json.Marshal(block)
	if err != nil {
		return
	}
	req := &p2p.MsgPto{
		Name: ip,
		Operation: Publish,
		Data: blockStr,
	}
	reqStr, err := json.Marshal(req)
	if err != nil || reqStr == nil {
		return
	}
	peerList := singleton.GetRouter().FetchPeers()
	if singleton.GetRouter().GetConnType() == p2p.ShortConnection {
		spreadShort(reqStr, peerList)
	}
}

// 同步等待和所有peer通信完毕
func spreadShort(reqStr []byte, peerList map[string]interface{}) {
	for k, _ := range peerList {
		wg.Add(1)
		go func(name string) {
			for reqStr != nil {
				b, err := singleton.Dispatch("yoghurt", reqStr)
				if err != nil {
					println("操作失败", err.Error())
					return
				}
				reqStr = nil
				reqStr, err = singleton.Handle(nil, b)
				fmt.Println(string(reqStr), err)
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
}

// TODO
func spreadPersistent(name string, resp []byte) {
}
