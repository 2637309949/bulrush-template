package models
import (
	"github.com/gin-gonic/gin"
)

// Inject -
func Inject(router *gin.RouterGroup) {
	user(router)
}
