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
	"github.com/2637309949/bulrush-template/cmd"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// HTTPTools defined  tools
type HTTPTools struct {
	engine *gin.Engine
}

func (t HTTPTools) handler(m, p string, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, nil)
	t.engine.ServeHTTP(w, req)
	h(w)
}

func (t HTTPTools) Get(p string, h func(w *httptest.ResponseRecorder)) {
	t.handler("GET", p, h)
}

func (t HTTPTools) Post(p string, h func(w *httptest.ResponseRecorder)) {
	t.handler("POST", p, h)
}

var httpTools *HTTPTools

func TestCache(t *testing.T) {
	httpTools.Get("/api/test/cache", func(w *httptest.ResponseRecorder) {
		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"message\":\"no token found\"}", w.Body.String())
	})
}

func TestMain(m *testing.M) {
	httpTools = &HTTPTools{}
	signal := make(chan struct{}, 0)
	app := cmd.New()
	go app.Run(func(httpProxy *gin.Engine, config *bulrush.Config) {
		httpTools.engine = httpProxy
		signal <- struct{}{}
	})
	<-signal
	os.Exit(m.Run())
}
