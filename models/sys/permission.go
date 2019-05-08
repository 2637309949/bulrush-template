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
	"github.com/globalsign/mgo/bson"
)

// Permission info
type Permission struct {
	models.Base `bson:",inline"`
	Code        string        `bson:"code" form:"code" json:"code" xml:"code"`
	Name        string        `bson:"name" form:"name" json:"name" xml:"name"`
	Pid         bson.ObjectId `bson:"pid" form:"pid" json:"pid" xml:"pid"`
	Type        string        `bson:"type" form:"type" json:"type" xml:"type"`
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "permission",
		"reflector": &Permission{},
	})
}

// RegisterPermission inject function
func RegisterPermission(r *gin.RouterGroup) {
	r.GET("/permission", addition.Mongo.Hooks.List("permission"))
	r.GET("/permission/:id", addition.Mongo.Hooks.One("permission"))
	r.POST("/permission", addition.Mongo.Hooks.Create("permission"))
	r.PUT("/permission", addition.Mongo.Hooks.Update("permission"))
	r.DELETE("/permission", addition.Mongo.Hooks.Delete("permission"))
}
