/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:29
 * @modify date 2019-01-16 20:49:29
 * @desc [description]
 */

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/models/sys"
	"github.com/gin-gonic/gin"
)

// Model register
var Model = bulrush.PNQuick(func(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(sys.RegisterUser)
	ri.Register(sys.RegisterPermission)
	ri.Register(sys.RegisterRole)
	ri.Register(sys.RegisterParam)
})
