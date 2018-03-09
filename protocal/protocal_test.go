package protocal

import (
	"encoding/json"
	"time"
	"fmt"
	"testing"

	p2p "github.com/Blockchain-CN/pheromones"
)

var (
	hello1 = p2p.MsgPto{
		Name:      "luda",
		Operation: ConnectReq,
	}
	hellomsg1 = MsgGreetingReq{
		Addr:    "127.0.0.1:12345",
		Account: 11900,
	}

	hello2 = p2p.MsgPto{
		Name:      "yoghurt",
		Operation: ConnectReq,
	}
	hellomsg2 = MsgGreetingReq{
		Addr:    "127.0.0.1:12346",
		Account: 11900,
	}
)

const (
	timeout = time.Millisecond * 100
)

func TestProtocal(t *testing.T) {
	Short()
}

func Short() {
	r1 := p2p.NewSRouter(timeout)
	p1 := NewProtocal("luda", r1, timeout)
	s1 := p2p.NewServer(p1, timeout)
	println("h1 监听 12345")
	go s1.ListenAndServe("127.0.0.1:12345")

	r2 := p2p.NewSRouter(timeout)
	p2 := NewProtocal("yoghurt", r2, timeout)
	s2 := p2p.NewServer(p2, timeout)
	println("h2 监听 12345")
	go s2.ListenAndServe("127.0.0.1:12346")

	r3 := p2p.NewSRouter(timeout)
	p3 := NewProtocal("diudiu", r3, timeout)
	s3 := p2p.NewServer(p3, timeout)
	println("h3 监听 12345")
	go s3.ListenAndServe("127.0.0.1:12347")

	time.Sleep(time.Second)
	p1.Add("yoghurt", "127.0.0.1:12346")
	j, _ := json.Marshal(hellomsg1)
	hello1.Data = j
	msg, _ := json.Marshal(hello1)
	for msg != nil {
		b, err := p1.Dispatch("yoghurt", msg)
		if err != nil {
			println("操作失败", err.Error())
			break
		}
		msg = nil
		msg, err = p1.Handle(nil, b)
		fmt.Println(string(msg), err)
	}
	fmt.Println("test1 done")

	j, _ = json.Marshal(hellomsg2)
	hello2.Data = j
	msg, _ = json.Marshal(hello2)
	for msg != nil {
		b, err := p2.Dispatch("luda", msg)
		if err != nil {
			println("操作失败", err.Error())
			break
		}
		msg = nil
		msg, err = p2.Handle(nil, b)
		fmt.Println(string(msg), err)
	}
	fmt.Println("test2 done")

	p2.Add("diudiu", "127.0.0.1:12347")
	j, _ = json.Marshal(hellomsg2)
	hello2.Data = j
	msg, _ = json.Marshal(hello2)
	b := p2.DispatchAll(msg)
	fmt.Println(b)
	fmt.Println("test3 done")

	p2.Delete("diudiu1")
	j, _ = json.Marshal(hellomsg2)
	hello2.Data = j
	msg, _ = json.Marshal(hello2)
	b = p2.DispatchAll(msg)
	fmt.Println(b)

	for {
		time.Sleep(time.Second)
	}
}
