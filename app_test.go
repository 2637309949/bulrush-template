// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var engine *gin.Engine

func handler(m, p string, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, nil)
	engine.ServeHTTP(w, req)
	h(w)
}

func Get(p string, h func(w *httptest.ResponseRecorder)) {
	handler("GET", p, h)
}

func Post(p string, h func(w *httptest.ResponseRecorder)) {
	handler("POST", p, h)
}

func TestCache(t *testing.T) {
	Get("/api/test/cache", func(w *httptest.ResponseRecorder) {
		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"message\":\"no token found\"}", w.Body.String())
	})
}

func TestMain(m *testing.M) {
	ok := make(chan struct{}, 0)
	gin.SetMode("release")
	app := app()
	go app.Run(func(httpProxy *gin.Engine, config *bulrush.Config) {
		engine = httpProxy
		ok <- struct{}{}
	})
	<-ok
	os.Exit(m.Run())
}
