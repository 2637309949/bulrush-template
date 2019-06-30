// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush_template/addition"
)

// User defined struct
type Role struct {
	Base
	Name string `bson:"name" form:"name" json:"name" xml:"name"`
	Type string `bson:"type" form:"type" json:"type" xml:"type"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "role",
	Reflector: &Role{},
})
