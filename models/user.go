package models

import (
	"github.com/2637309949/bulrush_template/utils"
	"github.com/gin-gonic/gin"
)

// User info
type User struct {
	Name string			`bson:"name"`
	Password string		`bson:"password"`
	Age int				`bson:"age"`
}

var manifest = map[string]interface{} {
	"db": "test",
	"name": "user",
	"collection": "user",
	"reflector": &User{},
}

// user inject function
func user(r *gin.RouterGroup) {
	utils.Mgo.Register(manifest)
	r.GET("/user", utils.Mgo.Hooks.List("user"))
	r.GET("/user/:id", utils.Mgo.Hooks.One("user"))
}
