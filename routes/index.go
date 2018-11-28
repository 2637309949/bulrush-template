package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
)

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	router, _ := injects["Router"].(*gin.RouterGroup)
	hello(router)
}

// Routes -
type Routes struct {
	bulrush.InjectGroup
}

// InjectRouter -
func (rs *Routes)InjectRouter(router *gin.RouterGroup) {
	hello(router)
}
