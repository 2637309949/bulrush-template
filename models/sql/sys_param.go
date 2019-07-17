// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
)

// Param defined struct
type Param struct {
	Model
	Code  string     `gorm:"comment:'编码';unique;not null"`
	Name  string     `gorm:"comment:'名称'"`
	Value []Property `gorm:"foreignkey:ParamID"`
}

// AddEnum defined add enum type
func (p *Param) AddEnum(model string, key string, value *[]Property) *Param {
	tx := addition.GORMExt.DB.Begin()
	enum := &Param{Model: PresetModel(), Code: "enum", Name: "枚举类型"}
	if err := tx.FirstOrCreate(enum, &Param{Code: "enum"}).Error; err != nil {
		tx.Rollback()
		addition.Logger.Error(err.Error())
		return p
	}
	for _, v := range *value {
		v.ParamID = enum.ID
		v.Category = model
		v.SubCategory = key
		v.Model = PresetModel()
		if err := tx.FirstOrCreate(&v, &Property{Category: v.Category, SubCategory: v.SubCategory, Key: v.Key}).Error; err != nil {
			tx.Rollback()
			addition.Logger.Error(err.Error())
			return p
		}
	}
	if err := tx.Commit().Error; err != nil {
		addition.Logger.Error(err.Error())
	}
	return p
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "Param",
	Reflector: &Param{},
})
