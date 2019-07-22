// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/models/sql"
	"github.com/2637309949/bulrush-template/plugins"
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /gorm/mock/login                模拟登录
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
func mockGormLogin(router *gin.RouterGroup) {
	router.GET("/gorm/mock/login", func(c *gin.Context) {
		user := addition.GORMExt.Var("User")
		if err := addition.GORMExt.DB.Find(user, map[string]interface{}{"name": "preset"}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		token, err := plugins.Identify.ObtainToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.SetCookie("accessToken", token.AccessToken, 60*60*24, "/", "", false, true)
		c.JSON(http.StatusOK, token.AccessToken)
	})
}

/**
 * @api {get} /gorm/mock/init                 GORM测试数据
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
func mockInit(router *gin.RouterGroup) {
	router.GET("/gorm/mock/init", func(c *gin.Context) {
		addition.GORMExt.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
		// 1. create user
		tx := addition.GORMExt.DB.Begin()
		tx.Exec("SET FOREIGN_KEY_CHECKS = 0;")
		tx.Exec("DROP DATABASE IF EXISTS bulrush;")
		tx.Exec("CREATE DATABASE bulrush;")
		tx.Exec("USE bulrush;")
		tx.DropTableIfExists(&sql.User{})
		tx.CreateTable(&sql.User{})
		first := &sql.User{
			Model: sql.PresetModel(),
			Name:  "L1211",
			Age:   23,
		}
		tx.Create(first)
		second := &sql.User{
			Model: sql.PresetModel(),
			Name:  "L1212",
			Age:   24,
		}
		tx.Create(second)

		// 2. create permission
		tx.DropTableIfExists(&sql.Permission{})
		tx.CreateTable(&sql.Permission{})
		tx.Create(&sql.Permission{
			Model: sql.PresetModel(),
			Code:  "ER5T12",
			Name:  "FINANCE MENU",
			Type:  "102",
		})
		tx.Create(&sql.Permission{
			Model: sql.PresetModel(),
			Code:  "ER5T15",
			Name:  "MENU",
			Type:  "101",
			Pid:   1,
		})

		// 3. create role
		tx.DropTableIfExists(&sql.Role{})
		tx.CreateTable(&sql.Role{})
		tx.Create(&sql.Role{
			Model: sql.PresetModel(),
			Name:  "FINANCE1",
			Type:  "101",
		})

		p1 := &sql.Permission{}
		p1.ID = 1
		tx.Create(&sql.Role{
			Model:       sql.PresetModel(),
			Name:        "FINANCE1",
			Type:        "101",
			Permissions: []*sql.Permission{p1},
		})

		tx.Exec("SET FOREIGN_KEY_CHECKS = 1;")
		err := tx.Commit().Error
		if err != nil {
			err = tx.Rollback().Error
		}
	})
}

// RegisterMock defined hello routes
func RegisterMock(ri *bulrush.ReverseInject) {
	ri.Register(mockGormLogin)
	ri.Register(mockInit)
}
