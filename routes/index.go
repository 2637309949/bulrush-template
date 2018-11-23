package routes

import (
	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	mongo, _ := injects["Mongo"].(*bulrush.MongoGroup)
	engine, _ := injects["Engine"].(*gin.Engine)
	router, _ := injects["Router"].(*gin.RouterGroup)
	hello(engine, router, mongo)
}