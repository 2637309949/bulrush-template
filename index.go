package main

import (
			 "github.com/gin-gonic/gin"
			 "github.com/2637309949/bulrush"
			 "github.com/2637309949/bulrush/utils"
	bMiddles "github.com/2637309949/bulrush/middles"
			 "github.com/2637309949/bulrush_template/routes"
			 "github.com/2637309949/bulrush_template/models"
			 "github.com/2637309949/bulrush_template/binds"
			 "github.com/2637309949/bulrush_template/middles"
   			 "errors"
			 "path"
			 "fmt"
			 "os"
)

// GINMODE app env
var GINMODE = utils.Some(os.Getenv("GIN_MODE"), "local")
// CONFIGPATH app config file path
var CONFIGPATH = path.Join(fmt.Sprintf("conf/%s.yaml", GINMODE))

func main() {
	iden := &bMiddles.Iden {
		ExpiresIn: 86400,
		Routes: bMiddles.RoutesGroup {
			ObtainTokenRoute:  "/api/v1/obtainToken",
			RevokeTokenRoute:  "/api/v1/revokeToken",
			RefleshTokenRoute: "/api/v1/fleshToken",
		},
		Auth: func(c *gin.Context) (interface{}, error) {
			var login binds.Login
			if err := c.ShouldBindJSON(&login); err != nil {
				return nil, err
			} else {
				if login.Password == "xx" && login.UserName == "xx" {
					return map[string] interface{} {
						"id":			"3e4r56u80a55",
						"username": 	login.UserName,
					}, nil
				}
				return nil, errors.New("User authentication failed")
			}
		},
		Tokens: bMiddles.TokensGroup {
			Save 	:func(interface{}) interface{} {
				return nil
			},
			Revoke  :func(interface{})  interface{} {
				return nil
			},
			Verify  :func(interface{})  interface{} {
				return nil
			},
		},
		IgnoreURLs: []interface{}{ `^/api/v1/ignore$` },
	}

	app := bulrush.New()
	app.Use(iden.Middle(app))
	app.Use(func(c *gin.Context) {
		fmt.Println("I am from bulrush middles!")
		c.Next()
	})
	app.LoadConfig(CONFIGPATH, utils.YAMLMode)
	// middles be about to load middles that from use func setting, and then this.
	app.Inject(middles.Inject, routes.Inject, models.Inject)
	app.Run()
}
