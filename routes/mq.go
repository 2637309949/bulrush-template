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

func mqHello(router *gin.RouterGroup, event events.EventEmmiter, q *mq.MQ) {
	q.Register("mqTest", func(mess mq.Message) error {
		fmt.Println(mess.Body)
		return nil
	})
	router.GET("/mqTest", func(c *gin.Context) {
		q.Push(mq.Message{
			Type: "mqTest",
			Body: map[string]interface{}{
				"message": "ok",
			},
		})
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}

// RegisterMq defined hello routes
func RegisterMq(router *gin.RouterGroup, event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(mqHello)
}
