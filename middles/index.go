package middles

import (
	"path"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/middles"
)

var override = middles.Override{}
var static = middles.Serve {
	URLPrefix: "/public",
	Fs: middles.LocalFile(path.Join("assets/public", ""), false),
}

// Middles -
type Middles struct {
	bulrush.InjectGroup
}

// Inject -
func (ms *Middles)Inject(injects map[string]interface{}) {
	override.Inject(injects)
	identity.Inject(injects)
	static.Inject(injects)
}

// InjectEngine -
func (ms *Middles)InjectEngine(engine *gin.Engine, router *gin.RouterGroup) {
	engine.Use(gin.Recovery())
	router.Use(gin.Recovery())
}
