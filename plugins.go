/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"net/http"

	"github.com/2637309949/bulrush"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/plugins"
	"github.com/gin-gonic/gin"
)

// appUsePlugins add plugin to application
func appUsePlugins(app bulrush.Bulrush) {
	app.Use(plugins.Limit)
	app.Use(plugins.Proxy)
	app.Use(plugins.Delivery)
	app.Use(plugins.Upload)
	app.Use(plugins.Logger)
	app.Use(plugins.Captcha)
	app.Use(plugins.Identify)
	app.Use(plugins.Role)
	app.Use(Model, Route)

	// mount models routers
	app.PostUse(addition.Mongo.AutoHook)
	// PNQuick Plugin init
	app.Use(bulrush.PNQuick(func(testInject string, role *role.Role, router *gin.RouterGroup) {
		router.GET("/bulrushApp", role.Can("r1,r2@p1,p3,p4;r4"), func(c *gin.Context) {
			addition.Logger.Info("from bulrushApp %s", "info")
			addition.Logger.Error("from bulrushApp %s", "error")
			c.JSON(http.StatusOK, gin.H{
				"message": testInject,
			})
		})
	}))
}
