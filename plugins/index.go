package plugins

import (
	"path"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush_template/binds"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/plugins"
	"github.com/2637309949/bulrush_template/utils"
)

// Override -
var override = &plugins.Override{}
// Delivery -
var delivery = &plugins.Delivery {
	URLPrefix: "/public",
	Fs: plugins.LocalFile(path.Join("assets/public", ""), false),
}
// Identify -
var identity = &plugins.Identify {
	ExpiresIn: 	86400,
	Auth: 	func(c *gin.Context) (interface{}, error) {
		var login binds.Login
		if err := c.ShouldBindJSON(&login); err != nil {
			return nil, err
		}
		if login.Password == "xx" && login.UserName == "xx" {
			return  map[string] interface{} {
						"id":			"3e4r56u80a55",
						"username": 	login.UserName,
					}, nil
		}
		return nil, errors.New("user authentication failed")
	},
	Routes: plugins.RoutesGroup {
		ObtainTokenRoute:  "/obtainToken",
		RevokeTokenRoute:  "/revokeToken",
		RefleshTokenRoute: "/refleshToken",
	},
	Tokens: plugins.TokensGroup {
				Save 	: utils.Rds.Hooks.SaveToken,
				Revoke  : utils.Rds.Hooks.RevokeToken,
				Find	: utils.Rds.Hooks.FindToken,
			},
	IgnoreURLs: []interface{}{ `^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$` },
}

// Middles -
type Middles struct {
	bulrush.InjectGroup
}

// Inject -
func (ms *Middles)Inject(injects map[string]interface{}) {
	override.Inject(injects)
	delivery.Inject(injects)
	identity.Inject(injects)
}
