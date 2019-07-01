// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"github.com/2637309949/bulrush"
	// import the docs package api compressor
	_ "github.com/2637309949/bulrush-template/docs"
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// APIDoc defined api docs
var APIDoc = bulrush.PNQuick(func(httpProxy *gin.Engine) {
	cfg := &swagger.Config{
		URL: "./docs/doc.json",
	}
	httpProxy.GET("/docs/*any", swagger.CustomWrapHandler(cfg, swaggerFiles.Handler))
})
