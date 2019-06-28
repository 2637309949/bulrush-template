// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

import (
	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

// OpenAPI register
// Make sure all models are initialized here
var OpenAPI = bulrush.PNQuick(func(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(RegisterHello)
})
