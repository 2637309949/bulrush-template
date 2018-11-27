package routes

import (
	"github.com/gin-gonic/gin"
)

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	router, _ := injects["Router"].(*gin.RouterGroup)
	hello(router)
}