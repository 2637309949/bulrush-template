package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
)

// Routes -
type Routes struct {
	bulrush.InjectGroup
}

// InjectRouter -
func (rs *Routes)InjectRouter(router *gin.RouterGroup) {
	hello(router)
}
