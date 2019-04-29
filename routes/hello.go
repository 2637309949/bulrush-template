/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:48
 * @modify date 2019-01-16 20:49:48
 * @desc [description]
 */

package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/2637309949/bulrush-template/models"
	"github.com/2637309949/bulrush-template/services"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/go-events"
)

func hello(router *gin.RouterGroup, event events.EventEmmiter) {
	event.On("hello", func(payload ...interface{}) {
		message := payload[0].(string)
		fmt.Println(message)
	})
	store := persistence.NewInMemoryStore(time.Second)

	/**
	     * @api {get} /ping 数据库测试
	     * @apiGroup Test
	     * @apiDescription 测试系统可用性
	     * @apiParam {String} accessToken        认证令牌
	     * @apiSuccessExample {json}             正常返回
	     * HTTP/1.1 200 OK
	     * {
		 *        "message":    "ok"
		 * }
	*/
	router.GET("/ping", func(c *gin.Context) {
		services.AddUsers([]interface{}{
			models.User{
				Name:     "double",
				Password: "111111",
				Age:      24,
			},
		})
		users := services.FindUsers(bson.M{"name": "double"})
		c.JSON(http.StatusOK, users)
	})

	/**
	     * @api {get} /chache 缓存测试
	     * @apiGroup Test
	     * @apiDescription 测试系统可用性
	     * @apiParam {String} accessToken        认证令牌
	     * @apiSuccessExample {json}             正常返回
	     * HTTP/1.1 200 OK
	     * {
		 *        "message":    "ok"
		 * }
	*/
	router.GET("/chache", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		fmt.Println("no chache")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}))
}
