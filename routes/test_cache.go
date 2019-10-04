// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"time"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

/**
 * @api {get} /test/chache                    缓存路由
 * @apiGroup Test
 * @apiDescription                            缓存路由
 * @apiSuccess        {Object}  Mess		  实体类
 * @apiSuccess        {String}  Mess.message  消息内容
 * @apiSuccessExample {json}                  正常返回
 * HTTP/1.1 200 OK
 * {
 *    "message": "hello"
 * }
**/
func chache(router *gin.RouterGroup, event events.EventEmmiter) {
	store := persistence.NewInMemoryStore(time.Second)
	router.GET("/test/cache", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": addition.I18N.I18NLocale("sys_hello", "hello"),
		})
	}))
}

// RegisterCache defined test routes
func RegisterCache(ri *bulrush.ReverseInject) {
	ri.Register(chache)
}
