// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"

	"github.com/2637309949/bulrush"
	mq "github.com/2637309949/bulrush-mq"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

/**
 * @api {get} /test/mq/hello                  队列路由
 * @apiGroup Test
 * @apiDescription                            队列路由
 * @apiSuccess        {Object}  Mess		  实体类
 * @apiSuccess        {String}  Mess.message  消息内容
 * @apiSuccessExample {json}                  正常返回
 * HTTP/1.1 200 OK
 * {
 *    "message": "ok"
 * }
**/
func mqHello(router *gin.RouterGroup, event events.EventEmmiter, q *mq.MQ) {
	router.GET("/test/mq/hello", func(c *gin.Context) {
		q.Push(mq.Message{
			Type: "mqTest",
			Body: map[string]interface{}{
				"message": "ok",
			},
		})
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	q.Register("mqTest", func(mess mq.Message) error {
		fmt.Println(mess.Body)
		return nil
	})
}

// RegisterMq defined hello routes
func RegisterMq(router *gin.RouterGroup, event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(mqHello)
}