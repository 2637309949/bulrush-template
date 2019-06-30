// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gormext"
	"github.com/2637309949/bulrush_template/addition"
)

// Permission defined struct
type Permission struct {
	Base
	Code       string      `form:"code" json:"code" xml:"code"`
	Name       string      `form:"name" json:"name" xml:"name"`
	Type       uint        `form:"type" json:"type" xml:"type"` // 1:'一级菜单', 2:'二级菜单', 3:'三级菜单', 4:'按钮', 5:'自定义'
	Pid        uint        `form:"pid" json:"pid" xml:"pid"`
	Permission *Permission `gorm:"foreignkey:id;association_foreignkey:pid" form:"permission" json:"permission" xml:"permission"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "permission",
	Reflector: &Permission{},
})
