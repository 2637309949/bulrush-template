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
	Code       string      `gorm:"comment:编码;"`
	Name       string      `gorm:"comment:名称;"`
	Type       uint        `gorm:"comment:类型;enum:一级菜单=1 二级菜单=2 三级菜单=3 按钮=4 自定义=5;"`
	Pid        uint        `gorm:"comment:父级ID;"`
	Permission *Permission `gorm:"foreignkey:id;association_foreignkey:pid"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "permission",
	Reflector: &Permission{},
})
