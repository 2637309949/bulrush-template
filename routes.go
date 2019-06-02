/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:54
 * @modify date 2019-01-16 20:49:54
 * @desc [description]
 */

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
	event.Emit("hello", "this is my payload to hello router")
})
