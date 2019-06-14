// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
)

// Role plugin init
var Role = &role.Role{
	RoleHandler: func(c *gin.Context, action string) bool {
		actions := role.TransformAction(action)
		addition.Logger.Debug("actions: %s\n", actions)
		return true
	},
}
