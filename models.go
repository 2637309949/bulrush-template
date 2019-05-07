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
	// Model -
	Model struct {
		bulrush.PNBase
	}
)

// Plugin -
func (model *Model) Plugin() bulrush.PNRet {
	return func(router *gin.RouterGroup) {
		sys.RegisterUser(router)
	}
}
