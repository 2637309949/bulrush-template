/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:40
 * @modify date 2019-01-16 20:49:40
 * @desc [description]
 */

package models

import (
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// User info
type User struct {
	Base     `bson:",inline"`
	Name     string          `bson:"name" form:"name" json:"name" xml:"name"`
	Password string          `bson:"password" form:"password" json:"password" xml:"password" `
	Age      int             `bson:"age" form:"age" json:"age" xml:"age"`
	IsActive bool            `bson:"isActive" form:"isActive" json:"isActive" xml:"isActive" `
	IsRepass bool            `bson:"isRepass" form:"isRepass" json:"isRepass" xml:"isRepass" `
	Avatar   string          `bson:"avatar" form:"avatar" json:"avatar" xml:"avatar" `
	Email    string          `bson:"email" form:"email" json:"email" xml:"email" `
	Roles    []bson.ObjectId `bson:"roles" form:"roles" json:"roles" xml:"roles" `
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "user",
		"reflector": &User{},
		"autoHook":  false,
	})
}

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	addition.Mongo.API.List(r, "user").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	})
	addition.Mongo.API.One(r, "user")
	addition.Mongo.API.Create(r, "user")
	addition.Mongo.API.Update(r, "user")
	addition.Mongo.API.Delete(r, "user")
}
