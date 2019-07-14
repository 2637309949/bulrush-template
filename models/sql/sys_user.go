// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
)

// User defined struct
type User struct {
	Model
	Name string `gorm:"comment:'名称'"`
	Age  uint   `gorm:"comment:'年龄'"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "User",
	Reflector: &User{},
	BanHook:   true,
	Opts: &gormext.Opts{
		RouteHooks: &gormext.RouteHooks{
			List: &gormext.ListHook{
				Pre: func(c *gin.Context) {
					addition.Logger.Info("User model pre hook")
				},
			},
		},
	},
})

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	addition.GORMExt.API.List(r, "User").Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.GORMExt.API.Feature("subUser").List(r, "User")
	addition.GORMExt.API.One(r, "User")
	addition.GORMExt.API.Create(r, "User")
	addition.GORMExt.API.Update(r, "User")
	addition.GORMExt.API.Delete(r, "User")
}
