// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
)

// Model common fields
type Model struct {
	gormext.Model
	Creator    *User `gorm:"foreignkey:id;association_foreignkey:creator_id"`
	CreatorID  uint  `gorm:"comment:'创建人ID'"`
	Modifier   *User `gorm:"foreignkey:id;association_foreignkey:modifier_id"`
	ModifierID uint  `gorm:"comment:'修改人ID'"`
	Deleter    *User `gorm:"foreignkey:id;association_foreignkey:deleter_id"`
	DeleterID  uint  `gorm:"comment:'删除人ID'"`
}