// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package models

import (
	"encoding/json"

	"github.com/Blockchain-CN/blockchain/common"
)

// TheChain BlockChain struct.
type TheChain struct {
	Chain []*Block `json:"chain"`
}

var singleChain *TheChain

func init() {
	singleChain = newChain()
	Genesis := GenerateBlock("0", "This is Genesis Block, Copyright Belong to Blockchain-CN", 0)
	singleChain.Chain = append(singleChain.Chain, Genesis)
}

func newChain() *TheChain {
	theChain := make([]*Block, 0)
	return &TheChain{theChain}
}

// FormatChain format received []byte to a blockchain object.
func FormatChain(b []byte) (*TheChain, error) {
	c := &TheChain{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, err
}

// AppendChain append a valid block to the chain's tail.
func AppendChain(b *Block) error {
	if !b.IsValid(GetChainTail()) {
		return common.Error(common.ErrInvalidBlock)
	}
	singleChain.Chain = append(singleChain.Chain, b)
	return nil
}

// FetchChain fetch the whole chain.
func FetchChain() *TheChain {
	return singleChain
}

// GetChainTail get the tail block of the chain.
func GetChainTail() *Block {
	return singleChain.Chain[GetChainLen()-1]
}

// GetChainLen get the chain's length.
func GetChainLen() int64 {
	return int64(len(singleChain.Chain))
}

// ReplaceChain replace the chain by a longer valid chain.
func ReplaceChain(c2 *TheChain) error {
	if int64(len(c2.Chain)) <= GetChainLen() {
		return common.Error(common.ErrInvalidBlock)
	}
	for i, b := range c2.Chain {
		if i == 0 {
			if *c2.Chain[i] != *singleChain.Chain[i] {
				return common.Error(common.ErrInvalidGenesisBlock)
			}
			continue
		}
		if !b.IsValid(c2.Chain[i-1]) {
			return common.Error(common.ErrInvalidBlock)
		}
	}
	singleChain.Chain = c2.Chain
	return nil
}
