// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User info
type User struct {
	Model    `bson:",inline"`
	Name     string          `bson:"name" br:"comment:'类别'"`
	Password string          `bson:"password" br:"comment:'类别'"`
	Age      uint            `bson:"age"`
	RoleIds  []bson.ObjectId `bson:"role_ids" br:"comment:'类别'"`
	Roles    *[]Role         `bson:"roles,omitempty" br:"ref(role,role_ids,_id)'"`
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
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
}).Init(func(ext *mgoext.Mongo) {
	Model := addition.MGOExt.Model("user")
	for _, key := range []string{"name"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := Model.EnsureIndex(index); err != nil {
			addition.Logger.Error(err.Error())
		}
	}
}).Init(func(ext *mgoext.Mongo) {
	// 添加预置账号, 用户存储预置数据
	Model := addition.MGOExt.Model("user")
	id := bson.ObjectIdHex("5d2fdc047dead1c7924b3a52")
	preset := &User{Model: PresetModel(), Name: "preset"}
	if err := Model.Find(bson.M{"name": "preset"}).One(preset); err == mgo.ErrNotFound {
		preset.ID = id
		if err := Model.Insert(preset); err != nil {
			addition.Logger.Error(err.Error())
		}
	}
	if preset.ID.Hex() != id.Hex() {
		if err := Model.RemoveId(preset.ID); err != nil {
			addition.Logger.Error(err.Error())
		} else {
			preset.ID = id
			preset.Password = "123456"
			if err := Model.Insert(preset); err != nil {
				addition.Logger.Error(err.Error())
			}
		}
	}
})

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup, role *role.Role) {
	addition.MGOExt.API.List(r, "user").Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.MGOExt.API.Feature("feature").List(r, "user")
	addition.MGOExt.API.One(r, "user", role.Can("r1,r2@p1,p3,p4;r4"))
	addition.MGOExt.API.Create(r, "user")
	addition.MGOExt.API.Update(r, "user")
	addition.MGOExt.API.Delete(r, "user")
}
