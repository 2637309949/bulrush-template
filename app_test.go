// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/2637309949/bulrush"
)

func setupRouter() *gin.Engine {
	var proxy *gin.Engine
	app := app()
	gin.SetMode("release")
	app.Run(func(httpProxy *gin.Engine, config *bulrush.Config) {
		proxy = httpProxy
	})
	return proxy
}

func TestCache(t *testing.T) {
	proxy := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/chache?accessToken=DEBUG", nil)
	proxy.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"ok\"}", w.Body.String())
}
