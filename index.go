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

	"github.com/2637309949/bulrush"
	delivery "github.com/2637309949/bulrush-delivery"
	identify "github.com/2637309949/bulrush-identify"
	logger "github.com/2637309949/bulrush-logger"
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
			if err := c.ShouldBindJSON(&login); err != nil {
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
			Save:   utils.Rds.Hooks.SaveToken,
			Revoke: utils.Rds.Hooks.RevokeToken,
			Find:   utils.Rds.Hooks.FindToken,
		},
		FakeURLs: []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
	})
	app.Use(&models.Model{}, &routes.Route{})
	app.Use(bulrush.PNQuick(func(testInject string, router *gin.RouterGroup) {
		router.GET("/bulrushApp", func(c *gin.Context) {
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
