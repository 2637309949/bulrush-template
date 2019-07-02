// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/2637309949/bulrush"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/models"
	"github.com/2637309949/bulrush-template/openapi"
	"github.com/2637309949/bulrush-template/plugins"
	"github.com/2637309949/bulrush-template/routes"
	"github.com/2637309949/bulrush-template/tasks"
	"github.com/gin-gonic/gin"
)

func appUsePlugins(app bulrush.Bulrush) {
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
	app.Use(models.Model, routes.Route, tasks.Task, openapi.OpenAPI)
	app.PostUse(addition.GORMExt, addition.MGOExt)
	app.PostUse()
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
