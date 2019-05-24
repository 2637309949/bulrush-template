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

// Role info
type Role struct {
	models.Base `bson:",inline"`
	Name        string          `bson:"name" form:"name" json:"name" xml:"name"`
	Type        string          `bson:"type" form:"type" json:"type" xml:"type"`
	Permissions []bson.ObjectId `bson:"permissions" form:"permissions" json:"permissions" xml:"permissions" `
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "role",
		"reflector": &Role{},
	})
}

// RegisterRole inject function
func RegisterRole(r *gin.RouterGroup) {
	addition.Mongo.API.List(r, "role")
	addition.Mongo.API.One(r, "role")
	addition.Mongo.API.Create(r, "role")
	addition.Mongo.API.Update(r, "role")
	addition.Mongo.API.Delete(r, "role")
}
