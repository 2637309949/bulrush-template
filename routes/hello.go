// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/models/nosql"
	"github.com/2637309949/bulrush-template/services"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/go-events"
)

/**
 * @api {get} /chache/xxx 测试缓存路由
 * @apiGroup Cache
 * @apiDescription xxxbbb
 * @apiSuccess {Object[]} profiles       List of user profiles.
 * @apiSuccess {Number}   profiles.age   Users age.
 * @apiSuccess {String}   profiles.image Avatar-Image.
 * @apiParam {String} accessToken    令牌
 * @apiParam {String} ids            顶级评分项ID, 如果多个就用用","分割
 * @apiParam {String} label          顶级评分项label, 如果多个就用用","分割
 * @apiSuccessExample {json}         正常返回
 * HTTP/1.1 200 OK
 * {
 *    "message": "ok"
 * }
**/
func cachePage(router *gin.RouterGroup, event events.EventEmmiter) {
	store := persistence.NewInMemoryStore(time.Second)
	router.GET("/chache", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		fmt.Println("no chache")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}))
}

func ping(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/mgoTest", func(c *gin.Context) {
		services.AddUsers([]interface{}{
			nosql.User{
				Name:     "double",
				Password: "111111",
				Age:      24,
				Model: nosql.Model{
					CreatorID:  bson.NewObjectId(),
					ModifierID: bson.NewObjectId(),
				},
			},
		})
		users := services.FindUsers(bson.M{"name": "double"})
		c.JSON(http.StatusOK, users)
	})
}

func onHello(event events.EventEmmiter) {
	event.On("hello", func(payload ...interface{}) {
		message := payload[0].(string)
		fmt.Println(message)
	})
}

func calltask(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/calltask", func(c *gin.Context) {
		event.Emit("reminderEmails", "hello 2019")
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}

// RegisterHello defined hello routes
func RegisterHello(router *gin.RouterGroup, event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(onHello)
	ri.Register(ping)
	ri.Register(cachePage)
	ri.Register(calltask)
}
