/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:54
 * @modify date 2019-01-16 20:49:54
 * @desc [description]
 */

package routes

import (
	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

type (
	// Route application routes
	Route struct {
		bulrush.PNBase
	}
)

// Plugin -
func (route *Route) Plugin() bulrush.PNRet {
	return func(router *gin.RouterGroup) {
		hello(router)
	}
}
