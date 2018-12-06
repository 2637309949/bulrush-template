package models

import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
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
func user(mongo *bulrush.MongoGroup, r *gin.RouterGroup) {
	mongo.Register(manifest)
	r.GET("/user", mongo.Hooks.List("user"))
	r.GET("/user/:id", mongo.Hooks.One("user"))
}
