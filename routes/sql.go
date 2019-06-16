// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"reflect"

	"github.com/2637309949/bulrush_template/addition"
	"github.com/2637309949/bulrush_template/models/sql"
	"github.com/gin-gonic/gin"
)

// RegisterSQL for routes
func RegisterSQL(router *gin.RouterGroup) {
	router.GET("/testsql", func(c *gin.Context) {
		addition.GORM.DB.AutoMigrate(&sql.Product{})
		addition.GORM.DB.Create(&sql.Product{Code: "L1212", Price: 1000})
		products := reflect.New(reflect.SliceOf(reflect.ValueOf(sql.Product{}).Type())).Interface()
		addition.GORM.DB.Find(products)
		c.JSON(http.StatusOK, products)
	})
}
