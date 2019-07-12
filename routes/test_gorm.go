// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"reflect"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/models/sql"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

/**
 * @api {get} /test/seq/users                     查询用户
 * @apiGroup Test
 * @apiDescription                                查询用户
 * @apiSuccess        {Object[]}  Users		      实体类数组
 * @apiSuccess        {String}    Users.Name      名称
 * @apiSuccess        {String}    Users.Password  密码
 * @apiSuccessExample {json}                      正常返回
 * HTTP/1.1 200 OK
 * {
 *     "ID": 1,
 *     "CreatedAt": "2019-07-12T20:59:57+08:00",
 *     "UpdatedAt": "2019-07-12T20:59:57+08:00",
 *     "DeletedAt": null,
 *     "Creator": null,
 *     "CreatorID": 0,
 *     "Modifier": null,
 *     "ModifierID": 0,
 *     "Deleter": null,
 *     "DeleterID": 0,
 *     "Name": "L1212",
 *     "Age": 23
 * }
**/
func testseq(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/test/seq/users", func(c *gin.Context) {
		addition.GORMExt.DB.AutoMigrate(&sql.User{})
		addition.GORMExt.DB.Create(&sql.User{Name: "L1212", Age: 23})
		products := reflect.New(reflect.SliceOf(reflect.ValueOf(sql.User{}).Type())).Interface()
		addition.GORMExt.DB.Find(products)
		c.JSON(http.StatusOK, products)
	})
}

// RegisterSeq for routes
func RegisterSeq(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(testseq)
}
