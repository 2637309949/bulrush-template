/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

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
	app := InitApp()
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