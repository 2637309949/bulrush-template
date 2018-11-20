package middles

import (
	// "2637309949/Bulrush/bulrush/middles"
	"github.com/gin-gonic/gin"
)

// Base as Function that accept some injection parameters
func Base (r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(middles.Override(r))
}