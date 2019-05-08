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

// Param info
type Param struct {
	models.Base `bson:",inline"`
	Code        string      `bson:"code" form:"code" json:"code" xml:"code"`
	Desc        string      `bson:"desc" form:"desc" json:"desc" xml:"desc"`
	Value       interface{} `bson:"value" form:"value" json:"value" xml:"value"`
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "param",
		"reflector": &Param{},
	})
}

// RegisterParam inject function
func RegisterParam(r *gin.RouterGroup) {
	r.GET("/param", addition.Mongo.Hooks.List("param"))
	r.GET("/param/:id", addition.Mongo.Hooks.One("param"))
	r.POST("/param", addition.Mongo.Hooks.Create("param"))
	r.PUT("/param", addition.Mongo.Hooks.Update("param"))
	r.DELETE("/param", addition.Mongo.Hooks.Delete("param"))
}
