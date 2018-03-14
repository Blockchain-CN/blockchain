// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package create

// CRequest request struct
type CRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// NewCRequestIDL ...
func NewCRequestIDL() *CRequest {
	return &CRequest{}
}

// CResponse response struct
type CResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

// NewCResponseIDL ...
func NewCResponseIDL() *CResponse {
	return &CResponse{}
}
