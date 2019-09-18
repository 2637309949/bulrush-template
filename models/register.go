// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/models/nosql"
	"github.com/2637309949/bulrush-template/models/sql"
	"github.com/gin-gonic/gin"
)

// Model register
// Make sure all models are initialized here
func Model(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(nosql.RegisterUser)
	ri.Register(nosql.RegisterRole)
	ri.Register(sql.RegisterUser)
}
