// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush_template/addition"

	"github.com/gin-gonic/gin"
)

// Product defined struct
type Product struct {
	Base
	Code  string `bson:"code" form:"code" json:"code" xml:"code"`
	Price uint   `bson:"price" form:"price" json:"price" xml:"price"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "product",
	Reflector: &Product{},
	BanHook:   true,
})

// RegisterProduct inject function
func RegisterProduct(r *gin.RouterGroup) {
	addition.GORMExt.API.List(r, "product").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.GORMExt.API.Feature("subProduct").List(r, "product")
	addition.GORMExt.API.One(r, "product")
	addition.GORMExt.API.Create(r, "product")
	addition.GORMExt.API.Update(r, "product")
	addition.GORMExt.API.Delete(r, "product")
}
