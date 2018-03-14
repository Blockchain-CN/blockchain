// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package common

// response返回错误码
const (
	// Success ...
	Success = 0

	// JoinPeerFail 添加peer失败
	JoinPeerFail = 1001
	// CreatBlockFail 添加block失败
	CreateBlockFain = 1002
)

// 内部错误码
const (
	// ErrInvalidBlock 非法区块
	ErrInvalidBlock = 10001
	// ErrInvalidGenesisBlock 非法创世区块
	ErrInvalidGenesisBlock = 10002

	// ErrInvalidChain 非法链
	ErrInvalidChain = 20001
)

// Error ...
type Error int

// Error ...
func (err Error) Error() string {
	return ErrMap[err]
}

// ErrMap 错误映射
var ErrMap = map[Error]string{
	Success:         "成功",
	JoinPeerFail:    "添加peer失败",
	CreateBlockFain: "添加block失败",

	ErrInvalidBlock:        "非法区块",
	ErrInvalidGenesisBlock: "非法创世区块",

	ErrInvalidChain: "非法链",
}
