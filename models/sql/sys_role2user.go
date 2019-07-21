// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
)

// Role2User defined m2m
type Role2User struct {
	Model
	UserID uint `gorm:"comment:'用户ID'"`
	RoleID uint `gorm:"comment:'角色ID'"`
}

// TableName defined TableName
func (Role2User) TableName() string {
	return "role2users"
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "Role2User",
	Reflector: &Role2User{},
})