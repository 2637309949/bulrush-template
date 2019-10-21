// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

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
func (r *Routes) testMQHello(router *gin.RouterGroup, event events.EventEmmiter, q *mq.MQ) {
	router.GET("/test/mq/hello", func(c *gin.Context) {
		q.Push(mq.Message{
			Type: "mqTest",
			Body: map[string]interface{}{
				"message": "ok",
			},
		})
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	q.Register("mqTest", func(mess *mq.Message) error {
		r.Logger.Info("%v", mess.Body)
		return nil
	})
}
