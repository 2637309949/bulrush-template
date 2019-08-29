// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"errors"

	identify "github.com/2637309949/bulrush-identify"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
)

// Identify plugin init
var Identify = identify.
	New().
	AddOptions(identify.FakeURLsOption([]string{`^/api/ignore$`, `^/api/gorm/mock`})).
	AddOptions(identify.FakeTokensOption([]string{})).
	AddOptions(identify.ModelOption(&identify.RedisModel{
		Redis: addition.Redis,
	})).
	Init(func(iden *identify.Identify) {
		iden.AddOptions(
			identify.AuthOption(func(ctx *gin.Context) (interface{}, error) {
				login := struct {
					UserName string `form:"username" json:"username" xml:"username" binding:"required"`
					Password string `form:"password" json:"password" xml:"password" binding:"required"`
					Type     string `form:"type" json:"type" xml:"type"`
				}{}
				// captcha := ctx.GetString("captcha")
				if err := ctx.ShouldBind(&login); err != nil {
					return nil, err
				}
				user := addition.GORMExt.Var("User")
				if err := addition.GORMExt.DB.Find(user, map[string]interface{}{"name": login.UserName, "password": login.Password}).Error; err != nil {
					return nil, errors.New("user authentication failed")
				}
				return user, nil
			}),
		)
	})
