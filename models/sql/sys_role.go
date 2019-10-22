// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
)

// Role defined struct
type Role struct {
	Model
	Code        string        `gorm:"comment:'编码';unique;not null"`
	Name        string        `gorm:"comment:'名称';unique;not null"`
	Type        string        `gorm:"comment:'类型';enum:'管理='101' 业务='102';not null"`
	Permissions []*Permission `gorm:"comment:'权限列表';many2many:role_permission;"`
}

var _ = GORMExt.Register(&gormext.Profile{
	Name:      "Role",
	Reflector: new(Role),
}).Init(func(ext *gormext.GORM) {
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
