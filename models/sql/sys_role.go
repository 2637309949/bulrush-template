// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
)

// Role defined struct
type Role struct {
	Model
	Code        string        `gorm:"comment:'编码';unique;not null"`
	Name        string        `gorm:"comment:'名称';unique;not null"`
	Type        string        `gorm:"comment:'类型';enum:'管理='101' 业务='102'"`
	Permissions []*Permission `gorm:"comment:'权限列表';many2many:role_permission;"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "Role",
	Reflector: &Role{},
}).Init(func(ext *gormext.GORM) {
	DB := addition.GORMExt.Model("Role")
	// 预置管理员角色
	preset := PresetModel()
	preset.ID = 101
	DB.Where("ID=101").Assign(&Role{Name: "管理员", Type: "101"}).FirstOrCreate(&Role{
		Model: preset,
		Name:  "管理员",
		Type:  "101",
	})
	(&Param{}).
		AddEnum("Role", "Type", &[]Property{
			Property{
				Key:   "管理",
				Value: "101",
			},
			Property{
				Key:   "业务",
				Value: "102",
			},
		})
})
