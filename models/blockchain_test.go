package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAppendChain(t *testing.T) {
	fmt.Println("the chain length is :", GetChainLen())

	b1 := GenerateBlock(GetChainTail().Hash, "Test_1", GetChainLen())
	err := AppendChain(b1)
	if err != nil {
		fmt.Println("Test_1 append fail, err :", err)
	}
	fmt.Println("the chain length is :", GetChainLen())

	b2 := GenerateBlock(GetChainTail().Hash, "Test_2", GetChainLen())
	err = AppendChain(b2)
	if err != nil {
		fmt.Println("Test_2 append fail, err :", err)
	}
	fmt.Println("the chain length is :", GetChainLen())

	b3 := GenerateBlock(GetChainTail().Hash, "Test_3", GetChainLen()+1)
	err = AppendChain(b3)
	if err != nil {
		fmt.Println("Test_3 append fail, err :", err)
	}
	fmt.Println("the chain length is :", GetChainLen())
}

func TestReplaceChain(t *testing.T) {
	fmt.Println("the chain length is :", GetChainLen())

	b1 := GenerateBlock(GetChainTail().Hash, "Test_1", GetChainLen())
	err := AppendChain(b1)
	if err != nil {
		fmt.Println("Test_1 append fail, err :", err)
	}
	fmt.Println("the chain length is :", GetChainLen())

	c := FetchChain()
	bc, _ := json.Marshal(c)
	c2, err := FormatChain(bc)
	if err != nil {
		fmt.Println("format fail")
	}
	if c2 != FetchChain() {
		fmt.Printf("format fail2, c2=%v||c=%v\n", c2, FetchChain())
	}

	b2 := GenerateBlock(GetChainTail().Hash, "Test_2", GetChainLen())
	err = AppendChain(b2)
	if err != nil {
		fmt.Println("Test_2 append fail, err :", err)
	}
	fmt.Println("the chain length is :", GetChainLen())

	err = ReplaceChain(c2)
	if err != nil {
		fmt.Println("replace fail, err:", err)
	}
	fmt.Println("the chain length is :", GetChainLen())
}
