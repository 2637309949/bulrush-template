// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
)

// Param defined struct
type Param struct {
	Model
	Code  string     `gorm:"comment:'编码';unique;not null"`
	Name  string     `gorm:"comment:'名称'"`
	Value []Property `gorm:"foreignkey:ParamID"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "Param",
	Reflector: &Param{},
	Opts: &gormext.Opts{
		RouteHooks: &gormext.RouteHooks{
			// overide global
			List: &gormext.ListHook{
				Cond: func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
					return cond
				},
			},
		},
	},
})

// AddEnum defined add enum type
func (p *Param) AddEnum(model string, key string, value *[]Property) *Param {
	DB := addition.GORMExt.Model("Param")
	tx := DB.Begin()
	enum := &Param{Model: DefaultModel(), Code: "enum", Name: "枚举类型"}
	if tx.Where(&Param{Code: "enum"}).Find(enum).RecordNotFound() {
		if err := tx.Save(&enum).Error; err != nil {
			tx.Rollback()
			addition.Logger.Error(err.Error())
			return p
		}
	}
	for _, v := range *value {
		v.ParamID = enum.ID
		v.Category = model
		v.SubCategory = key
		if tx.Where(&Property{Category: v.Category, SubCategory: v.SubCategory, Key: v.Key}).Find(&v).RecordNotFound() {
			v.Model = DefaultModel()
			if err := tx.Save(&v).Error; err != nil {
				tx.Rollback()
				addition.Logger.Error(err.Error())
				return p
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		addition.Logger.Error(err.Error())
	}
	return p
}
