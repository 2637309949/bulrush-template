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
	Creator   *User `gorm:"foreignkey:id;association_foreignkey:creator_id"`
	UpdatorID uint  `gorm:"comment:'修改人ID'"`
	Updator   *User `gorm:"foreignkey:id;association_foreignkey:updator_id"`
	DeleterID uint  `gorm:"comment:'删除人ID'"`
	Deleter   *User `gorm:"foreignkey:id;association_foreignkey:deleter_id"`
}

// PresetModel defined Preset User
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
