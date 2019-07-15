// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type test struct {
	Info1 string `bson:"indo1" br:"comment:'测试1'"`
}

// User info
type User struct {
	Model    `bson:",inline"`
	Name     string          `bson:"name" br:"comment:'类别'"`
	Password string          `bson:"password" br:"comment:'类别'"`
	Age      uint            `bson:"age"`
	RoleIds  []bson.ObjectId `bson:"role_ids" br:"comment:'类别'"`
	Roles    *[]Role         `bson:"roles,omitempty" br:"ref(role,role_ids,_id)'"`
	Test1    *test           `bson:"test1" br:"comment:'测试1'"`
	Test2    *test
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	Name:      "User",
	Reflector: &User{},
	BanHook:   true,
	Opts: &mgoext.Opts{
		RouteHooks: &mgoext.RouteHooks{
			List: &mgoext.ListHook{
				Pre: func(c *gin.Context) {
					addition.Logger.Info("user before")
				},
			},
		},
	},
})

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup, role *role.Role) {
	addition.MGOExt.API.List(r, "User").Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.MGOExt.API.Feature("feature").List(r, "User")
	addition.MGOExt.API.One(r, "User", role.Can("r1,r2@p1,p3,p4;r4"))
	addition.MGOExt.API.Create(r, "User")
	addition.MGOExt.API.Update(r, "User")
	addition.MGOExt.API.Delete(r, "User")
}
