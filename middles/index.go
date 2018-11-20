package middles

import (
	"github.com/gin-gonic/gin"
)

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	engine, ok := injects["Engine"].(*gin.Engine)
	if ok {
		Base(engine)
	}
}