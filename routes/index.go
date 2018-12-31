package routes

import (
	"github.com/gin-gonic/gin"
)


// Inject -
func Inject(router *gin.RouterGroup) {
	hello(router)
}
