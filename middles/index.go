package middles

import (
	"path"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush/middles"
)

var static = middles.Serve {
	URLPrefix: "/public",
	Fs: middles.LocalFile(path.Join("assets/public", ""), false),
}

// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	e, _ := injects["Engine"].(*gin.Engine);
	r, _ := injects["Router"].(*gin.RouterGroup);
	
	e.Use(gin.Recovery())
	r.Use(middles.Override(e))
	identity.Inject(injects)
	static.Inject(injects)
}