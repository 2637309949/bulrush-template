// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"reflect"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush_template/addition"
	"github.com/2637309949/bulrush_template/models/sql"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

// @Summary MGO测试
// @Description MGO测试
// @Tags GORM
// @Accept mpfd
// @Produce json
// @Param accessToken query string true "令牌"
// @Success 200 {string} json "{"message": "ok"}"
// @Failure 400 {string} json "{"message": "failure"}"
// @Router /testsql [get]
func testsql(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/testsql", func(c *gin.Context) {
		addition.GORMExt.DB.AutoMigrate(&sql.Product{})
		addition.GORMExt.DB.Create(&sql.Product{Code: "L1212", Price: 1000})
		products := reflect.New(reflect.SliceOf(reflect.ValueOf(sql.Product{}).Type())).Interface()
		addition.GORMExt.DB.Find(products)
		c.JSON(http.StatusOK, products)
	})
}

// RegisterSQL for routes
func RegisterSQL(router *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(testsql)
}
