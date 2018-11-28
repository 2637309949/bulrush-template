package middles

import (
	"path"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/middles"
)

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
	engine, _ := injects["Engine"].(*gin.Engine);
	router, _ := injects["Router"].(*gin.RouterGroup);
	router.Use(middles.Override(engine))
	identity.Inject(injects)
	static.Inject(injects)
}

// InjectEngine -
func (ms *Middles)InjectEngine(engine *gin.Engine) {
	engine.Use(gin.Recovery())
}
