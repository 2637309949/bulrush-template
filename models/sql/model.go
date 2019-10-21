// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"time"

	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/utils"
)

// GORMExt gorm ext
var GORMExt = addition.GORMExt

// Logger logger
var Logger = addition.Logger

// Model defined common field
type Model struct {
	ID        uint       `gorm:"comment:'模型ID';primary_key"`
	CreatorID uint       `gorm:"comment:'创建人ID'"`
	Creator   *User      `gorm:"foreignkey:id;association_foreignkey:CreatorID"`
	CreatedAt *time.Time `gorm:"comment:'创建时间'"`

	UpdatorID uint       `gorm:"comment:'修改人ID'"`
	Updator   *User      `gorm:"foreignkey:id;association_foreignkey:UpdatorID"`
	UpdatedAt *time.Time `gorm:"comment:'更新时间'"`

	DeleterID uint       `gorm:"comment:'删除人ID'"`
	Deleter   *User      `gorm:"foreignkey:id;association_foreignkey:DeleterID"`
	DeletedAt *time.Time `sql:"index" gorm:"comment:'删除时间'"`
}

// DefaultModel defined DefaultModel
func DefaultModel() Model {
	return Model{CreatedAt: utils.Now(), UpdatedAt: utils.Now()}
}
