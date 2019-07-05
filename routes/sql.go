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

func testsql(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/testsql", func(c *gin.Context) {
		addition.GORMExt.DB.AutoMigrate(&sql.User{})
		addition.GORMExt.DB.Create(&sql.User{Name: "L1212", Age: 23})
		products := reflect.New(reflect.SliceOf(reflect.ValueOf(sql.User{}).Type())).Interface()
		addition.GORMExt.DB.Find(products)
		c.JSON(http.StatusOK, products)
	})
}

// RegisterSQL for routes
func RegisterSQL(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(testsql)
}
