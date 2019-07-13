// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/plugins"
	"github.com/gin-gonic/gin"
)

func mockLogin(router *gin.RouterGroup) {
	router.GET("/mgo/mock/login", func(c *gin.Context) {
		User := addition.MGOExt.Model("user")
		user := addition.MGOExt.Var("user")
		if err := User.Find(map[string]interface{}{"name": "double"}).One(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		token, err := plugins.Identify.ObtainToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.SetCookie("accessToken", token.AccessToken, 60*60*24, "/", "*", false, true)
		c.JSON(http.StatusOK, token.AccessToken)
	})
	router.GET("/gorm/mock/login", func(c *gin.Context) {
		user := addition.GORMExt.Var("user")
		if err := addition.GORMExt.DB.Find(user, map[string]interface{}{"name": "L1212"}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		token, err := plugins.Identify.ObtainToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.SetCookie("accessToken", token.AccessToken, 60*60*24, "/", "*", false, true)
		c.JSON(http.StatusOK, token.AccessToken)
	})
}

// RegisterMock defined hello routes
func RegisterMock(ri *bulrush.ReverseInject) {
	ri.Register(mockLogin)
}
