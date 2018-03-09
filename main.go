// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package main

import (
	"fmt"
	"os"

	"github.com/Blockchain-CN/blockchain/server"
	"github.com/Blockchain-CN/blockchain/common"
	pto "github.com/Blockchain-CN/blockchain/protocal"
)

func main() {
	// init protocal
	pto.InitPto("127.0.0.1:12346", common.P2PTimeOut)

	// call this func will block current goroutine
	if err := server.Serve(); err != nil {
		printAndDie(err)
		return
	}
}

func printAndDie(err error) {
	fmt.Fprintf(os.Stderr, "init failed, err:%s", err)
	os.Exit(-1)
}
