package models
import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
)

// Model -
type Model struct {
	bulrush.InjectGroup
}

// // Inject -
// func (ml *Model)Inject(injects map[string]interface{}) {
// 	mongo, _ := injects["Mongo"].(*bulrush.MongoGroup)
// 	router, _ := injects["Router"].(*gin.RouterGroup)
// 	user(mongo, router)
// }

// InjectMongoAndRouter -
func (ml *Model)InjectMongoAndRouter(mongo *bulrush.MongoGroup, router *gin.RouterGroup) {
	user(mongo, router)
}
