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
	Init(func(iden *identify.Identify) {
		iden.Auth = func(ctx *gin.Context) (interface{}, error) {
			login := struct {
				UserName string `form:"username" json:"username" xml:"username" binding:"required"`
				Password string `form:"password" json:"password" xml:"password" binding:"required"`
				Type     string `form:"type" json:"type" xml:"type"`
			}{}
			// captcha := ctx.GetString("captcha")
			if err := ctx.ShouldBind(&login); err != nil {
				return nil, err
			}
			if login.Password == "xx" && login.UserName == "xx" {
				return map[string]interface{}{
					"id":       "3e4r56u80a55",
					"username": login.UserName,
				}, nil
			}
			return nil, errors.New("user authentication failed")
		}
		iden.Model = &identify.RedisModel{
			Redis: addition.Redis,
		}
		iden.FakeTokens = []string{"DEBUG"}
		iden.FakeURLs = []string{`^/api/v1/ignore$`, `^/api/v1/ptest$`, `^/api/v1/gorm/mock`}
	})
