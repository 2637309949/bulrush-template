// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"github.com/2637309949/bulrush_template/addition"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Product defined struct
type Product struct {
	gorm.Model
	Code  string `bson:"code" form:"code" json:"code" xml:"code"`
	Price uint   `bson:"price" form:"price" json:"price" xml:"price"`
}

// Register model
func init() {
	addition.GORM.Register(map[string]interface{}{
		"db":        "test",
		"name":      "product",
		"reflector": &Product{},
		"autoHook":  false,
	})
}

// RegisterProduct inject function
func RegisterProduct(r *gin.RouterGroup) {
	addition.GORM.API.List(r, "product").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.GORM.API.One(r, "product")
	addition.GORM.API.Create(r, "product")
	addition.GORM.API.Update(r, "product")
	addition.GORM.API.Delete(r, "product")
}
