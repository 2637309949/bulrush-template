// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"net/http"

	"github.com/2637309949/bulrush"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/2637309949/bulrush-template/grpc"
	"github.com/2637309949/bulrush-template/models"
	"github.com/2637309949/bulrush-template/openapi"
	"github.com/2637309949/bulrush-template/plugins"
	"github.com/2637309949/bulrush-template/routes"
	"github.com/2637309949/bulrush-template/tasks"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

// New defined Bulrush
func New() bulrush.Bulrush {
	app := bulrush.Default()
	app.Config(conf.CPath)
	app.PreUse(addition.I18N)
	app.Use(plugins.Limit)
	app.Use(plugins.Proxy)
	app.Use(plugins.Delivery)
	app.Use(plugins.Upload)
	app.Use(plugins.Logger)
	app.Use(plugins.Captcha)
	app.Use(plugins.Identify)
	app.Use(plugins.Role)
	app.Use(plugins.OpenAPI)
	app.Use(plugins.MQ)
	app.Use(models.Init, routes.NewRoutes, grpc.Init, tasks.Init, openapi.Init)
	app.PostUse(addition.GORMExt, addition.MGOExt)
	app.PostUse(addition.APIDoc)
	app.Inject("bulrushApp")
	app.Use(func(testInject string, role *role.Role, router *gin.RouterGroup) {
		router.GET("/bulrushApp", role.Can("r1,r2@p1,p3,p4;r4"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": testInject,
			})
		})
	})
	app.Use(func(test string) {
		addition.Logger.Info("inject test= %s", test)
	}, func(event events.EventEmmiter) {
		event.On(bulrush.EventsRunning, func(message ...interface{}) {
			addition.Logger.Info("message from EventsRunning%v", message)
		})
	})

	return app
}

// Execute defined run Bulrush
func Execute() {
	New().Run()
}
