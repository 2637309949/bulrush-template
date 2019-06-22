// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush_template/addition"
	_ "github.com/2637309949/bulrush_template/docs"
	"github.com/2637309949/bulrush_template/plugins"
	"github.com/kataras/go-events"
)

// @title bulrush-template api
// @version 1.0
// @description bulrush web library template
// @termsOfService https://github.com/2637309949/bulrush

// @contact.name bulrush-template
// @contact.url https://github.com/2637309949/bulrush
// @contact.email 2637309949@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	app := InitApp()
	app.Use(plugins.APIDoc)
	app.Use(bulrush.PNQuick(func(event events.EventEmmiter) {
		event.On(bulrush.EventSysBulrushPluginRunImmediately, func(message ...interface{}) {
			addition.Logger.Info("EventSysBulrushPluginRunImmediately %v", message)
		})
	}))
	app.RunImmediately()
}
