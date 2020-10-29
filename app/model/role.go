package model

import "van-api/types"

type Role struct {
	ID         uint64
	Keyid      string
	Name       types.JSON
	Resource   string
	Acl        string
	Note       string
	Status     bool
	CreateTime uint64
	UpdateTime uint64
}
