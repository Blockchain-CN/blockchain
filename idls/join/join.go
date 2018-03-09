// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package join

// JRequest request struct
type JRequest struct {
	PeerAddr string `json:"peer_addr"`
}

// NewJRequest ...
func NewJRequest() *JRequest {
	return &JRequest{}
}

// JResponse response struct
type JResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

// NewJResponse ...
func NewJResponse() *JResponse {
	return &JResponse{}
}
