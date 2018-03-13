// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package protocal

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	p2p "github.com/Blockchain-CN/pheromones"
)

const (
	// 连接请求
	ConnectReq = "connectreq"
	// 获取一个
	GetReq = "getreq"
	// 批量获取
	FetchReq = "fetchreq"
	// 同步更新
	NoticeReq = "noticereq"

	// 连接请求返回
	ConnectResp = "connectresp"
	// 获取一个返回
	GetResp = "getresp"
	// 批量获取返回
	FetchResp = "fetchresp"
	// 同步更新返回
	NoticeResp = "noticeresp"

	// 未知操作
	UnknownOp = "unknownop"

	defultByte = 10240
)

type MsgGreetingReq struct {
	Addr    string `json:"add"`
	Account int    `json:"account"`
}

type Protocal2 struct {
	HostName string
	Router   p2p.Router
	to       time.Duration
}

func NewProtocal2(name string, r p2p.Router, to time.Duration) *Protocal2 {
	return &Protocal2{name, r, to}
}

func (p *Protocal2) Handle(c net.Conn, msg []byte) ([]byte, error) {
	cType := p.Router.GetConnType()
	req := &p2p.MsgPto{}
	resp := &p2p.MsgPto{}
	err := json.Unmarshal(msg, req)
	if err != nil {
		resp.Name = p.HostName
		resp.Operation = UnknownOp
		ret, _ := json.Marshal(resp)
		return ret, p2p.Error(p2p.ErrMismatchProtocalReq)
	}
	resp.Name = p.HostName
	switch req.Operation {
	case ConnectReq:
		subReq := &MsgGreetingReq{}
		err := json.Unmarshal(req.Data, subReq)
		if err != nil {
			return nil, p2p.Error(p2p.ErrMismatchProtocalResp)
		}
		if cType == p2p.ShortConnection {
			err = p.Router.AddRoute(req.Name, subReq.Addr)
		} else {
			if p.Router.AddRoute(req.Name, c) == nil {
				go p.IOLoop(c)
			}
		}
		if err != nil {
			fmt.Printf("@%s@report: %s operation from @%s@ err=%s\n", p.HostName, req.Operation, req.Name, err)
		}
		resp.Operation = ConnectResp
	case GetReq:
		resp.Operation = GetResp
	case FetchReq:
		resp.Operation = FetchResp
	case NoticeReq:
		resp.Operation = NoticeResp
	case ConnectResp:
		resp.Operation = GetReq
	case GetResp:
		resp.Operation = FetchReq
	case FetchResp:
		resp.Operation = NoticeReq
	case NoticeResp:
		fmt.Printf("@%s@report: %s operation from @%s@ finished\n", p.HostName, req.Operation, req.Name)
		return nil, nil
	default:
		resp.Operation = UnknownOp
	}
	ret, err := json.Marshal(resp)
	fmt.Printf("@%s@report: %s operation from @%s@ succeed\n", p.HostName, req.Operation, req.Name)
	return ret, nil
}

// 长连接的话，需要在加入路由的时刻起携程 循环监控
func (p *Protocal2) IOLoop(c net.Conn) {
	fmt.Printf("@%s@report,开启长连接监听: localhost=%s||remotehost=%s\n", p.HostName, c.LocalAddr(), c.RemoteAddr())
	for {
		msg, err := p.read(c)
		if err != nil {
			c.Close()
			return
		}
		fmt.Printf("长连接收到信息, localhost=%s||remotehost=%s||msg=%s\n", c.LocalAddr(), c.RemoteAddr(), string(msg))
		resp, err := p.Handle(c, msg)
		if err != nil || resp == nil {
			fmt.Printf("结束此次会话, localconn=%s||remoteconn=%s||resp=%s||err=%s\n", c.LocalAddr(), c.RemoteAddr(), resp, err)
			continue
		}
		c.SetWriteDeadline(time.Now().Add(p.to))
		_, err = c.Write(resp)
		if err != nil {
			return
		}
		fmt.Printf("长连接发送信息, localconn=%s||remoteconn=%s||msg=%s\n", c.LocalAddr(), c.RemoteAddr(), resp)
	}
}

func (p *Protocal2) read(r io.Reader) ([]byte, error) {
	buf := make([]byte, defultByte)
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	// read读出来的是[]byte("abcdefg"+0x00)，带一个结束符，需要去掉
	return buf[:n], nil
}

func (p *Protocal2) Add(name string, addr string) error {
	if p.Router.GetConnType() == p2p.ShortConnection {
		return p.Router.AddRoute(name, addr)
	}
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	if p.Router.AddRoute(name, c) == nil {
		go p.IOLoop(c)
	}
	return err
}

func (p *Protocal2) GetRouter() p2p.Router {
	return p.Router
}

func (p *Protocal2) DispatchAll(msg []byte) map[string][]byte {
	return p.Router.DispatchAll(msg)
}

func (p *Protocal2) Dispatch(name string, msg []byte) ([]byte, error) {
	return p.Router.Dispatch(name, msg)
}

func (p *Protocal2) Delete(name string) error {
	return p.Router.Delete(name)
}
