// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
)

// Permission defined struct
type Permission struct {
	Base
	Code       string      `gorm:"comment:'编码';"`
	Name       string      `gorm:"comment:'名称';"`
	Type       uint        `gorm:"comment:'类型';enum(1:'一级菜单', 2:'二级菜单', 3:'三级菜单', 4:'按钮', 5:'自定义');"`
	Pid        uint        `gorm:"comment:'父级ID';"`
	Permission *Permission `gorm:"foreignkey:id;association_foreignkey:pid"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "permission",
	Reflector: &Permission{},
})
