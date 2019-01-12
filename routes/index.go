package routes

import (
	"github.com/gin-gonic/gin"
)


// Plugin -
func Plugin(router *gin.RouterGroup) {
	hello(router)
}
