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
	"github.com/2637309949/bulrush-template/routes/sys"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

type (
	// Route for all routes register
	Route struct {
		bulrush.PNBase
	}
)

// Plugin for all routes register
func (route *Route) Plugin() bulrush.PNRet {
	return func(router *gin.RouterGroup, event events.EventEmmiter) {
		sys.RegisterHello(router, event)
		event.Emit("hello", "this is my payload to hello router")
	}
}
