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

// Role info
type Role struct {
	Model       `bson:",inline"`
	Name        string          `bson:"name" br:"comment:'编码'"`
	Type        string          `bson:"type" br:"comment:'类别'"`
	Permissions []bson.ObjectId `bson:"permissions" br:"comment:'权限ID'"`
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	DB:        "test",
	Name:      "role",
	Reflector: &Role{},
	BanHook:   true,
	Opts: &mgoext.Opts{
		RouteHooks: &mgoext.RouteHooks{
			List: &mgoext.ListHook{
				Auth: func(c *gin.Context) bool {
					addition.Logger.Info("Role model auth hook")
					return true
				},
			},
		},
	},
})

// RegisterRole inject function
func RegisterRole(r *gin.RouterGroup) {
	addition.MGOExt.API.ALL(r, "role").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	})
	addition.MGOExt.API.Feature("subRole").Feature("subRole2").List(r, "role").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Auth(func(c *gin.Context) bool {
		return false
	})
}
