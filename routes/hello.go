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
	"github.com/globalsign/mgo/bson"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush_template/services"
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush_template/utils"
)

func hello (r *gin.RouterGroup) {
	r.GET("/ping", func (c *gin.Context) {
		utils.Logger("pingtest...")
		services.AddUsers([]interface {} {
			models.User{
				Name: "double",
				Password: "111111",
				Age: 24,
			},
		})
		users := services.FindUsers(bson.M{"name": "double"})
		c.JSON(http.StatusOK, gin.H{
			"data": 	users,
			"errcode": 	nil,
			"errmsg": 	nil,
		})
	})

    /**
     * @api {get} /hello 检验Token有效性
     * @apiGroup Test
     * @apiDescription 测试系统可用性
     * @apiParam {String} accessToken        认证令牌
     * @apiSuccessExample {json}             正常返回
     * HTTP/1.1 200 OK
     * {
     *        "errcode": null,
     *        "errmsg":  null,
	 *        "data":    "ok"
	 * }
     */
	r.GET("/hello", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": 	"ok",
			"errcode": 	nil,
			"errmsg": 	nil,
		})
	})

	r.POST("/ptest", func(c *gin.Context) {
		var json map[string]interface{}
		if c.BindJSON(&json) == nil {
			fmt.Print(json)
		}
	})
}