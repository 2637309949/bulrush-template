/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:40
 * @modify date 2019-01-16 20:49:40
 * @desc [description]
 */

package sys

import (
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/models"
	"github.com/gin-gonic/gin"
)

// User info
type User struct {
	models.Base `bson:",inline"`
	Name        string `bson:"name" form:"name" json:"name" xml:"name"`
	Password    string `bson:"password" form:"password" json:"password" xml:"password" `
	Age         int    `bson:"age" form:"age" json:"age" xml:"age"`
	IsActive    bool   `bson:"isActive" form:"isActive" json:"isActive" xml:"isActive" `
	IsRepass    bool   `bson:"isRepass" form:"isRepass" json:"isRepass" xml:"isRepass" `
	Avatar      string `bson:"avatar" form:"avatar" json:"avatar" xml:"avatar" `
	Email       string `bson:"email" form:"email" json:"email" xml:"email" `
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "user",
		"reflector": &User{},
	})
}

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	r.GET("/user", addition.Mongo.Hooks.List("user"))
	r.GET("/user/:id", addition.Mongo.Hooks.One("user"))
	r.POST("/user", addition.Mongo.Hooks.Create("user"))
	r.PUT("/user", addition.Mongo.Hooks.Update("user"))
	r.DELETE("/user", addition.Mongo.Hooks.Delete("user"))
}
