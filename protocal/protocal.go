package protocal

import (
	"time"
	"encoding/json"
	"fmt"
	"io"
	"net"

	p2p "github.com/Blockchain-CN/pheromones"
	"github.com/Blockchain-CN/blockchain/models"
	"github.com/Blockchain-CN/blockchain/common"
)

type command string

const (
	// RequireBlock 请求最新block
	RequireBlock command = "require a block"

	// Publish 发布最新block
	DeliveryBlock command = "delivery a block"

	// DeliveryChain 发送整条链
	DeliveryChain command = "delivery the block"

	// RequireChain 请求整条链
	RequireChain command = "require the block"

	// 未知命令
	UnknownCmd command= "unknown cmd"

	defultByte = 10240
)

type Protocal struct {
	HostName string
	Router   p2p.Router
	to       time.Duration
}

func NewProtocal(name string, r p2p.Router, to time.Duration) *Protocal {
	return &Protocal{name, r, to}
}

func (p *Protocal) Handle(c net.Conn, msg []byte) ([]byte, error) {
	cType := p.Router.GetConnType()
	req := &p2p.MsgPto{}
	resp := &p2p.MsgPto{}
	err := json.Unmarshal(msg, req)
	if err != nil {
		resp.Name = p.HostName
		resp.Operation = UnknownCmd
		ret, _ := json.Marshal(resp)
		return ret, p2p.Error(p2p.ErrMismatchProtocalReq)
	}
	resp.Name = p.HostName
	switch req.Operation {
	case RequireBlock:
		if cType == p2p.ShortConnection {
			err = p.Router.AddRoute(req.Name, req.Name)
		} else {
			if p.Router.AddRoute(req.Name, c) == nil {
				go p.IOLoop(c)
			}
		}
		c, _ := json.Marshal(models.GetChainTail())
		resp.Operation = DeliveryBlock
		resp.Data = c
	case DeliveryBlock:
		block , err := models.FormatBlock(req.Data)
		if err != nil {
			return nil, p2p.Error(p2p.ErrMismatchProtocalResp)
		}
		// if the block's index is shorter or invalidate
		tailBlock := models.GetChainTail()
		if !block.IsTempValid() || block.Index <= tailBlock.Index {
			return nil, common.Error(common.ErrInvalidBlock)
		}
		// if the block can append to the tail
		if block.IsValid(tailBlock) {
			models.AppendChain(block)
			// 并需要向外广播
			go p.spreads(block)
			return nil, nil
		}
		// if the block's index is longer
		resp.Operation = RequireChain
	case RequireChain:
		c, _ := json.Marshal(models.FetchChain())
		resp.Operation = DeliveryChain
		resp.Data = c
	case DeliveryChain:
		chain, err := models.FormatChain(req.Data)
		if err != nil {
			return nil, p2p.Error(p2p.ErrMismatchProtocalResp)
		}
		err = models.ReplaceChain(chain)
		if err != nil {
			return nil, common.Error(common.ErrInvalidChain)
		}
		// 向外广播 models.GetChainTail()
		go p.spreads(models.GetChainTail())
		return nil, nil
	default:
		fmt.Printf("@%s@report: %s operation from @%s@ finished\n", p.HostName, req.Operation, req.Name)
		return nil, nil
	}
	ret, err := json.Marshal(resp)
	fmt.Printf("@%s@report: %s operation from @%s@ succeed\n", p.HostName, req.Operation, req.Name)
	return ret, nil
}

// 长连接的话，需要在加入路由的时刻起携程 循环监控
func (p *Protocal) IOLoop(c net.Conn) {
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

func (p *Protocal) read(r io.Reader) ([]byte, error) {
	buf := make([]byte, defultByte)
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	// read读出来的是[]byte("abcdefg"+0x00)，带一个结束符，需要去掉
	return buf[:n], nil
}

func (p *Protocal) Add(name string, addr string) error {
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

func (p *Protocal) GetRouter() p2p.Router {
	return p.Router
}

func (p *Protocal) DispatchAll(msg []byte) map[string][]byte {
	return p.Router.DispatchAll(msg)
}

func (p *Protocal) Dispatch(name string, msg []byte) ([]byte, error) {
	return p.Router.Dispatch(name, msg)
}

func (p *Protocal) Delete(name string) error {
	return p.Router.Delete(name)
}

// spread the latest block to all peers
func (p *Protocal) spreads(block *models.Block) {
	blockStr, err := json.Marshal(block)
	if err != nil {
		return
	}
	req := &p2p.MsgPto{
		Name: ip,
		Operation: DeliveryBlock,
		Data: blockStr,
	}
	reqStr, err := json.Marshal(req)
	if err != nil || reqStr == nil {
		return
	}
	peerList := p.GetRouter().FetchPeers()
	if p.GetRouter().GetConnType() == p2p.ShortConnection {
		p.spreadShort(reqStr, peerList)
	}
}

// 同步等待和所有peer通信完毕
func (p *Protocal) spreadShort(reqStr []byte, peerList map[string]interface{}) {
	for _, v := range peerList {
		wg.Add(1)
		go func(addr string) {
			for reqStr != nil {
				b, err := p.Dispatch(addr, reqStr)
				if err != nil {
					println("操作失败", err.Error())
					return
				}
				reqStr = nil
				reqStr, err = p.Handle(nil, b)
				fmt.Println(string(reqStr), err)
			}
			wg.Done()
		}(v.(string))
	}
	wg.Wait()
}

// TODO
func (p *Protocal) spreadPersistent(name string, resp []byte) {
}