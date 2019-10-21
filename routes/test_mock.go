// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

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
func (r *Routes) testMockGormLogin(router *gin.RouterGroup) {
	router.GET("/gorm/mock/login", func(c *gin.Context) {
		user := &sql.User{
			Name:  "test",
			Model: sql.DefaultModel(),
		}
		if r.GORMExt.DB.Where(map[string]interface{}{"name": "test"}).Find(user).RecordNotFound() {
			r.GORMExt.DB.Create(user)
		}
		user.CreatorID = user.ID
		user.UpdatorID = user.ID
		r.GORMExt.DB.Save(user)
		token, err := plugins.Identify.ObtainToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": addition.I18N.I18NLocale(err.Error(), err.Error()),
			})
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
func (r *Routes) testMockInit(router *gin.RouterGroup) {
	router.GET("/gorm/mock/init", func(c *gin.Context) {
		r.GORMExt.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
		// 1. create user
		tx := r.GORMExt.DB.Begin()
		tx.Exec("SET FOREIGN_KEY_CHECKS = 0;")
		tx.Exec("DROP DATABASE IF EXISTS bulrush;")
		tx.Exec("CREATE DATABASE bulrush;")
		tx.Exec("USE bulrush;")
		tx.DropTableIfExists(&sql.User{})
		tx.CreateTable(&sql.User{})
		first := &sql.User{
			Model: sql.DefaultModel(),
			Name:  "L1211",
			Age:   23,
		}
		tx.Create(first)
		first.CreatorID = first.ID
		first.UpdatorID = first.ID
		tx.Save(first)

		second := &sql.User{
			Model: sql.DefaultModel(),
			Name:  "L1212",
			Age:   24,
		}
		tx.Create(second)
		second.CreatorID = second.ID
		second.UpdatorID = second.ID
		tx.Save(second)

		model := sql.DefaultModel()
		model.CreatorID = first.ID
		model.UpdatorID = first.ID

		// 2. create permission
		tx.DropTableIfExists(&sql.Permission{})
		tx.CreateTable(&sql.Permission{})
		tx.Create(&sql.Permission{
			Model: model,
			Code:  "ER5T12",
			Name:  "FINANCE MENU",
			Type:  "102",
		})
		tx.Create(&sql.Permission{
			Model: model,
			Code:  "ER5T15",
			Name:  "MENU",
			Type:  "101",
			Pid:   1,
		})

		// 3. create role
		tx.DropTableIfExists(&sql.Role{})
		tx.CreateTable(&sql.Role{})
		tx.Create(&sql.Role{
			Model: model,
			Name:  "FINANCE1",
			Type:  "101",
		})

		p1 := &sql.Permission{}
		tx.Create(&sql.Role{
			Model:       model,
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
