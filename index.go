/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/2637309949/bulrush-template/addition"

	"github.com/2637309949/bulrush"
	delivery "github.com/2637309949/bulrush-delivery"
	identify "github.com/2637309949/bulrush-identify"
	logger "github.com/2637309949/bulrush-logger"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/binds"
	"github.com/2637309949/bulrush-template/models"
	"github.com/2637309949/bulrush-template/routes"
	"github.com/2637309949/bulrush-template/utils"
	upload "github.com/2637309949/bulrush-upload"
	"github.com/gin-gonic/gin"
)

// GOMODE APP ENV
var GOMODE = utils.Some(os.Getenv("GO_MODE"), "local")

// CONFIGPATH PATH
var CONFIGPATH = path.Join(".", fmt.Sprintf("conf/%s.yaml", GOMODE))

func main() {
	app := bulrush.Default()
	app.Config(CONFIGPATH)
	app.Inject("bulrushApp")
	app.Use(
		&delivery.Delivery{
			URLPrefix: "/public",
			Path:      path.Join("assets/public", ""),
		},
		&upload.Upload{
			URLPrefix: "/public/upload",
			AssetPath: path.Join("assets/public/upload", ""),
		},
		&logger.LoggerWriter{},
	)
	app.Use(&identify.Identify{
		Auth: func(c *gin.Context) (interface{}, error) {
			var login binds.Login
			if err := c.ShouldBind(&login); err != nil {
				return nil, err
			}
			if login.Password == "xx" && login.UserName == "xx" {
				return map[string]interface{}{
					"id":       "3e4r56u80a55",
					"username": login.UserName,
				}, nil
			}
			return nil, errors.New("user authentication failed")
		},
		Tokens: identify.TokensGroup{
			Save:   addition.Redis.Hooks.SaveToken,
			Revoke: addition.Redis.Hooks.RevokeToken,
			Find:   addition.Redis.Hooks.FindToken,
		},
		FakeTokens: []interface{}{"DEBUG"},
		FakeURLs:   []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
	})
	app.Use(&role.Role{
		FailureHandler: func(c *gin.Context, action string) {
		},
		RoleHandler: func(c *gin.Context, action string) bool {
			actions := role.TransformAction(action)
			fmt.Printf("actions: %s\n", actions)
			return true
		},
	})
	app.Use(&models.Model{}, &routes.Route{})
	app.Use(bulrush.PNQuick(func(testInject string, role *role.Role, router *gin.RouterGroup) {
		router.GET("/bulrushApp", role.Can("r1,r2@p1,p3,p4;r4"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": testInject,
			})
		})
	}))
	app.Run(func(err error, config *bulrush.Config) {
		if err != nil {
			panic(err)
		} else {
			name := config.GetString("name", "")
			port := config.GetString("port", "")
			fmt.Println("================================")
			fmt.Printf("App: %s\n", name)
			fmt.Printf("Listen on %s\n", port)
			fmt.Println("================================")
		}
	})
}
