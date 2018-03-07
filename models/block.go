// Copyright 2018 Lothar . All rights reserved.
// https://github.com/Blockchain-CN

package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	dhash "github.com/Blockchain-CN/sha256"

	"github.com/Blockchain-CN/blockchain/common"
)

// Block struct.
type Block struct {
	PVHash    string `json:"pv_hash"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Index     int64  `json:"index"`
	Nonce     int64  `json:"nonce"`
	Hash      string `json:"hash"`
}

// FormatBlock Format received []byte to a block object.
func FormatBlock(b []byte) (*Block, error) {
	block := &Block{}
	err := json.Unmarshal(b, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GenerateBlock Generate a new block, it takes sometime and can be stopped by using the following function.
// hash = PVHash+Timestamp+Data+n+Nonce.
func GenerateBlock(pvHash, data string, index int64) *Block {
	var metaData string
	time := time.Now().UnixNano()
	tStr := strconv.FormatInt(time, 10)
	nStr := strconv.FormatInt(index, 10)
	metaData = pvHash + tStr + data + nStr
	hash, nonce := dhash.HashwithDifficulty([]byte(metaData), common.HashDifficulty)
	return &Block{
		PVHash:    pvHash,
		Timestamp: time,
		Data:      data,
		Index:     index,
		Nonce:     nonce,
		Hash:      fmt.Sprint(hash),
	}
}

// Interupt stop calculating hash for the block.
func (b *Block) Interupt() bool {
	return dhash.StopHash()
}

// IsValid return if the block is legal.
func (b *Block) IsValid(pvb *Block) bool {
	var metaData string
	if b.PVHash != pvb.Hash || (pvb.Index+1) != b.Index {
		return false
	}
	tStr := strconv.FormatInt(b.Timestamp, 10)
	nStr := strconv.FormatInt(b.Index, 10)
	noStr := strconv.FormatInt(b.Nonce, 10)
	metaData = b.PVHash + tStr + b.Data + nStr
	return dhash.Verification(append([]byte(metaData), []byte(noStr)...), b.Hash)
}
