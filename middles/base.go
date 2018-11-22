package middles

import (
	"github.com/2637309949/bulrush/middles"
	"github.com/gin-gonic/gin"
)

func base (r *gin.RouterGroup, g *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middles.Override(g))
}