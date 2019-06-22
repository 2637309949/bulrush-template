// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/routes"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

// Route for all routes register
// Make sure all routes are initialized here
var Route = bulrush.PNQuick(func(router *gin.RouterGroup, event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(routes.RegisterHello)
	ri.Register(routes.RegisterSQL)
	ri.Register(routes.RegisterMq)
	event.Emit("hello", "this is my payload to hello router")
})
