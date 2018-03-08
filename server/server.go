// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package server

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Blockchain-CN/blockchain/server/create"
	"github.com/Blockchain-CN/blockchain/server/join"
	"github.com/Blockchain-CN/httpsvr"
)

// Serve ...
func Serve() error {
	s := httpsvr.New("127.0.0.1:10024",
		httpsvr.SetReadTimeout(time.Millisecond*200),
		httpsvr.SetWriteTimeout(time.Millisecond*200),
		httpsvr.SetMaxAccess(100),
	)
	go GracefulExit(s)
	s.AddRoute("POST", "/blockchain/create", &create.CController{})
	s.AddRoute("POST", "/blockchain/join", &join.JController{})
	return s.Serve()
}

// GracefulExit 优雅退出
func GracefulExit(svr *httpsvr.Server) {
	sigc := make(chan os.Signal, 0)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	<-sigc
	println("closing agent...")
	svr.GracefulExit()
	println("agent closed.")
	os.Exit(0)
}
