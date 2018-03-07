// Copyright 2018 Blockchain-CN . All rights reserved.
// https://github.com/Blockchain-CN

package models

import (
	"path"

	"github.com/Blockchain-CN/keygen"

	"github.com/Blockchain-CN/blockchain/common"
)

// User struct.
type User struct {
	Name    string
	Path    string
	Public  string
	Private string
}

// Login allow user to login and get their key.
func Login(name string) (*User, error) {
	uPath := keygen.GetUserPath(name)
	pvKeyPath := path.Join(uPath, "private.pem")
	pbKeyPath := path.Join(uPath, "public.pem")
	pv, errv := keygen.GetKeyMd5(pvKeyPath)
	pb, errb := keygen.GetKeyMd5(pbKeyPath)
	if errv == nil && errb == nil {
		return &User{
			Name:    name,
			Path:    uPath,
			Public:  pb,
			Private: pv}, nil
	}

	if err := keygen.GenRsaKey(common.RSADefaultLenth, name); err != nil {
		return nil, err
	}
	pv, err := keygen.GetKeyMd5(pvKeyPath)
	if err != nil {
		return nil, err
	}
	pb, err = keygen.GetKeyMd5(pbKeyPath)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:    name,
		Path:    uPath,
		Public:  pb,
		Private: pv}, nil
}
