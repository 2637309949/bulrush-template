package models
import (
	"github.com/gin-gonic/gin"
)

// Plugin -
func Plugin(router *gin.RouterGroup) {
	user(router)
}
