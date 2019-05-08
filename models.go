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

type (
	// Model for all model register
	Model struct {
		bulrush.PNBase
	}
)

// Plugin for all model register
func (model *Model) Plugin() bulrush.PNRet {
	return func(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
		ri.Register(sys.RegisterUser)
	}
}
