// Copyright 2018 Lothar . All rights reserved.
// https://github.com/Blockchain-CN

package common

const (
	// ErrInvalidBlock 非法区块
	ErrInvalidBlock = 1001
	// ErrInvalidGenesisBlock 非法创世区块
	ErrInvalidGenesisBlock = 1002

	// ErrInvalidChain 非法链
	ErrInvalidChain = 2001
)

// Error ...
type Error int

// Error ...
func (err Error) Error() string {
	return errMap[err]
}

var errMap = map[Error]string{
	ErrInvalidBlock:        "非法区块",
	ErrInvalidGenesisBlock: "非法创世区块",

	ErrInvalidChain: "非法链",
}
