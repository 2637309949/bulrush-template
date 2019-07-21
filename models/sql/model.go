// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"time"

	gormext "github.com/2637309949/bulrush-addition/gorm"
)

// Model defined common field
type Model struct {
	gormext.Model
	CreatorID uint  `gorm:"comment:'创建人ID'"`
	Creator   *User `gorm:"foreignkey:id;association_foreignkey:CreatorID"`
	UpdatorID uint  `gorm:"comment:'修改人ID'"`
	Updator   *User `gorm:"foreignkey:id;association_foreignkey:UpdatorID"`
	DeleterID uint  `gorm:"comment:'删除人ID'"`
	Deleter   *User `gorm:"foreignkey:id;association_foreignkey:DeleterID"`
}

// PresetModel defined Preset Model
// 系统内置数据时的默认参数
func PresetModel() Model {
	now := time.Now()
	return Model{
		Model: gormext.Model{
			UpdatedAt: &now,
			CreatedAt: &now,
		},
		CreatorID: 101,
		UpdatorID: 101,
	}
}

// PresetUser defined Preset User
// 系统内置数据时的默认参数
func PresetUser() User {
	model := PresetModel()
	model.ID = 101
	return User{
		Model:    model,
		Name:     "preset",
		Password: "123456",
	}
}

// PresetRole defined Preset Role
// 系统内置数据时的默认参数
func PresetRole() Role {
	model := PresetModel()
	model.ID = 101
	return Role{
		Model: model,
		Name:  "管理员",
		Code:  "8JN6GH",
		Type:  "101",
	}
}
