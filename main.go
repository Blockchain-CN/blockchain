// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package main

import (
	"fmt"
	"os"
	"errors"

	"github.com/Blockchain-CN/blockchain/common"
	pto "github.com/Blockchain-CN/blockchain/protocal"
	"github.com/Blockchain-CN/blockchain/server"
	"net"
)

func main() {
	ip := GetIP()
	if ip == "" {
		printAndDie(errors.New("Unable to get a avilable ip"))
	}

	// init protocal
	pto.InitPto(ip+":12346", common.P2PTimeOut)

	// call this func will block current goroutine
	if err := server.Serve(); err != nil {
		printAndDie(err)
		return
	}
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}

func printAndDie(err error) {
	fmt.Fprintf(os.Stderr, "init failed, err:%s", err)
	os.Exit(-1)
}
