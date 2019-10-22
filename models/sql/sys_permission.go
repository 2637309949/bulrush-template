// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
)

// Permission defined struct
type Permission struct {
	Model
	Code       string      `gorm:"comment:'编码';unique;not null"`
	Name       string      `gorm:"comment:'名称';unique;not null"`
	Pid        uint        `gorm:"comment:'父级ID'"`
	Type       string      `gorm:"comment:'类型';enum:'一级菜单='101' 二级菜单='102' 三级菜单='103' 按钮='104' 自定义='105''"`
	Holder     string      `gorm:"comment:'权限类型';enum:'系统='101' 用户='102'"`
	Permission *Permission `gorm:"foreignkey:id;association_foreignkey:p_id"`
}

var _ = GORMExt.Register(&gormext.Profile{
	Name:      "Permission",
	Reflector: new(Permission),
}).Init(func(ext *gormext.GORM) {
	// Generate front-end parameters
	(&Param{}).
		AddEnum("Permission", "Type", &[]Property{
			Property{
				Key:   "一级菜单",
				Value: "101",
			},
			Property{
				Key:   "二级菜单",
				Value: "102",
			},
			Property{
				Key:   "三级菜单",
				Value: "103",
			},
			Property{
				Key:   "按钮",
				Value: "104",
			},
			Property{
				Key:   "自定义",
				Value: "105",
			},
		}).
		AddEnum("Permission", "Holder", &[]Property{
			Property{
				Key:   "系统",
				Value: "101",
			},
			Property{
				Key:   "用户",
				Value: "102",
			},
		})
})
