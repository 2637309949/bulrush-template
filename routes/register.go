// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

// Route for all routes register
// Make sure all routes are initialized here
var Route = func(router *gin.RouterGroup, event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(RegisterMgo)
	ri.Register(RegisterCache)
	ri.Register(RegisterSeq)
	ri.Register(RegisterMq)
	ri.Register(RegisterEvent)
	event.Emit("hello", "this is my payload to hello router")
}
