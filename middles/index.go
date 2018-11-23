package middles

import (
	"github.com/gin-gonic/gin"
)

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	engine, _ := injects["Engine"].(*gin.Engine);
	router, _ := injects["Router"].(*gin.RouterGroup);
	base(router, engine)
	iden(injects)
	static(injects)
}