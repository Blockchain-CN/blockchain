// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package show

// JRequest request struct
type SRequest struct {
	Chain	bool	`json:"chain"`
	Peer 	bool	`json:"peer"`
}

// NewJRequest ...
func NewJRequest() *SRequest {
	return &SRequest{}
}

// JResponse response struct
type SResponse struct {
	Chain  interface{}	`json:"chain"`
	Peer   interface{}	`json:"peer"`
}

// NewJResponse ...
func NewJResponse() *SResponse {
	return &SResponse{}
}
