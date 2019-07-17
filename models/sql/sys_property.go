// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
)

// Property defined struct
type Property struct {
	Model
	ParamID     uint   `gorm:"comment:'Param外键'"`
	Category    string `gorm:"comment:'类别'"`
	SubCategory string `gorm:"comment:'子类别'"`
	Key         string `gorm:"comment:'编码'"`
	Value       string `gorm:"comment:'属性值'"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "Property",
	Reflector: &Property{},
})
