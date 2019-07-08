// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// User info
type User struct {
	Model    `bson:",inline"`
	Name     string          `br:"up" bson:"name" br:"comment:'类别'"`
	Password string          `br:"up" bson:"password" br:"comment:'类别'"`
	Age      uint            `br:"up" bson:"age"`
	Roles    []bson.ObjectId `br:"ref:role;up" bson:"roles" br:"comment:'类别'"`
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	DB:        "test",
	Name:      "user",
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
func RegisterUser(r *gin.RouterGroup) {
	addition.MGOExt.API.List(r, "user").Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.MGOExt.API.Feature("feature").List(r, "user")
	addition.MGOExt.API.One(r, "user")
	addition.MGOExt.API.Create(r, "user")
	addition.MGOExt.API.Update(r, "user")
	addition.MGOExt.API.Delete(r, "user")
}
