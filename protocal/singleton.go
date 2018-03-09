// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package protocal

import (
	p2p "github.com/Blockchain-CN/pheromones"

	"time"
)

var singleton *Protocal

func InitPto(addr string, to time.Duration) {
	r1 := p2p.NewSRouter(to)
	p1 := NewProtocal(addr, r1, to)
	s1 := p2p.NewServer(p1, to)
	singleton = p1
	println("P2P Servering on ", addr)
	go s1.ListenAndServe(addr)
}

func AddPeer(addr string) error {
	return singleton.Add(addr, addr)
}