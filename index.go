package main

import (
	"github.com/2637309949/bulrush_template/binds"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
	bUtils "github.com/2637309949/bulrush/utils"
	"github.com/2637309949/bulrush_template/routes"
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush_template/utils"
	"github.com/2637309949/bulrush/plugins"
	"path"
	"fmt"
	"os"
)

// GOMODE APP ENV
var GOMODE 	= bUtils.Some(os.Getenv("GO_MODE"), "local")
// CONFIGPATH PATH
var CONFIGPATH  = path.Join(".", fmt.Sprintf("conf/%s.yaml", GOMODE))
// Delivery -
var delivery = &plugins.Delivery {
	URLPrefix: "/public",
	Fs: plugins.LocalFile(path.Join("assets/public", ""), false),
}
// Override -
var override = &plugins.Override{}
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

func main() {
	gin.SetMode(gin.DebugMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Printf("%5v %9v\n", httpMethod, absolutePath)
	}
	app := bulrush.Default()
	app.Config(CONFIGPATH)
	app.Inject("bulrushApp")
	app.Use(override.Inject, delivery.Inject)
	app.Use(identity.Inject)
	app.Use(models.Inject, routes.Inject)
	app.Use(func(iStr string, router *gin.RouterGroup) {
		router.GET("/bulrushApp", func (c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": 	iStr,
				"errcode": 	nil,
				"errmsg": 	nil,
			})
		})
	})
	app.Run()
}
