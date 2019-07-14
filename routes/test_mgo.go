// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/models/nosql"
	"github.com/2637309949/bulrush-template/services"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/go-events"
)

/**
 * @api {get} /test/mgo/adduser                   添加用户
 * @apiGroup Test
 * @apiDescription                                添加用户
 * @apiSuccess        {Object[]}  Users		      实体类数组
 * @apiSuccess        {String}    Users.Name      名称
 * @apiSuccess        {String}    Users.Password  密码
 * @apiParamExample {json} Request-Example:
 *     {
 *       "cond": { "age": { $gte: 1 } },
 *       "range": "page",
 *       "page": 1,
 *       "size": 10,
 *       "project": "_id,name",
 *       "preload": "Creator",
 *       "order": "-_created",
 *     }
 * @apiSuccessExample {json}                      正常返回
 * HTTP/1.1 200 OK
 * {
 *     "ID": "5d064cc213756cdb0f662607",
 *     "Created": null,
 *     "Modified": null,
 *     "Deleted": null,
 *     "CreatorID": "5d064cc2e8cd7d3029885465",
 *     "Creator": null,
 *     "ModifierID": "5d064cc2e8cd7d3029885466",
 *     "Modifier": null,
 *     "DeleterID": "",
 *     "Deleter": null,
 *     "Name": "double",
 *     "Password": "111111",
 *     "Age": 24,
 *     "RoleIds": null,
 *     "Roles": [],
 *     "Test1": null,
 *     "Test2": null
 * }
**/
func adduser(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/test/mgo/adduser", func(c *gin.Context) {
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

// RegisterMgo defined test routes
func RegisterMgo(ri *bulrush.ReverseInject) {
	ri.Register(adduser)
}
