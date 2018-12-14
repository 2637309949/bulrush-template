package models
import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
)

// Model -
type Model struct {
	bulrush.InjectGroup
}

// InjectMongoAndRouter -
func (ml *Model)InjectMongoAndRouter(router *gin.RouterGroup) {
	user(router)
}
