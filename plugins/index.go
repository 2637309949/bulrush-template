package plugins

import (
	"path"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/plugins"
	"github.com/2637309949/bulrush_template/utils"
)

// Recovery -
type Recovery struct {
}

// Override -
var override = &plugins.Override{}
// Delivery -
var delivery = &plugins.Delivery {
	URLPrefix: "/public",
	Fs: plugins.LocalFile(path.Join("assets/public", ""), false),
}
// Iden -
var identity = &plugins.Iden {
	ExpiresIn: 	86400,
	Routes: plugins.RoutesGroup {
				ObtainTokenRoute:  "/obtainToken",
				RevokeTokenRoute:  "/revokeToken",
				RefleshTokenRoute: "/refleshToken",
			},
	Auth: 	utils.Auth,
	Tokens: plugins.TokensGroup {
				Save 	: utils.Save,
				Revoke  : utils.Revoke,
				Find	: utils.Find,
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
	identity.Inject(injects)
	delivery.Inject(injects)
}
