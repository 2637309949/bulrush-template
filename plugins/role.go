/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import (
	"fmt"

	role "github.com/2637309949/bulrush-role"
	"github.com/gin-gonic/gin"
)

// Role plugin init
var Role = &role.Role{
	FailureHandler: func(c *gin.Context, action string) {
	},
	RoleHandler: func(c *gin.Context, action string) bool {
		actions := role.TransformAction(action)
		fmt.Printf("actions: %s\n", actions)
		return true
	},
}
