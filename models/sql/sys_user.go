// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"time"

	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"
)

// User defined struct
type User struct {
	Model
	Name     string                 `gorm:"comment:'名称';unique;not null"`
	Password string                 `gorm:"comment:'密码';not null"`
	Age      uint                   `gorm:"comment:'年龄'"`
	Birthday *time.Time             `gorm:"comment:'生日'"`
	Mobile   string                 `gorm:"comment:'手机'"`
	Email    null.String            `gorm:"comment:'邮箱'"`
	Roles    []*Role                `gorm:"comment:'角色列表';many2many:role2users;"`
	Ignored  *struct{ Name string } `gorm:"-"`
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
	}).RouteHooks(&gormext.RouteHooks{
		// Override global config, never query only by own
		List: &gormext.ListHook{
			Cond: func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
				return cond
			},
		},
	})
	addition.GORMExt.API.Feature("subUser").List(r, "User")
	addition.GORMExt.API.One(r, "User")
	addition.GORMExt.API.Create(r, "User")
	addition.GORMExt.API.Update(r, "User")
	addition.GORMExt.API.Delete(r, "User")
}
