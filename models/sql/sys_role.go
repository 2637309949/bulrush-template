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
	Name        string        `gorm:"comment:'name of role';"`
	Type        uint          `gorm:"comment:'type of role';"`
	Permissions []*Permission `gorm:"comment:'permissions in role';many2many:role_permission;"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "role",
	Reflector: &Role{},
})
