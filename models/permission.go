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

// Permission info
type Permission struct {
	Base `bson:",inline"`
	Code string        `bson:"code" form:"code" json:"code" xml:"code"`
	Name string        `bson:"name" form:"name" json:"name" xml:"name"`
	Pid  bson.ObjectId `bson:"pid" form:"pid" json:"pid" xml:"pid"`
	Type string        `bson:"type" form:"type" json:"type" xml:"type"`
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "permission",
		"reflector": &Permission{},
		"autoHook":  false,
	})
}

// RegisterPermission inject function
func RegisterPermission(r *gin.RouterGroup) {
	addition.Mongo.API.ALL(r, "permission")
}
