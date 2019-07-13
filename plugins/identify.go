// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"errors"

	identify "github.com/2637309949/bulrush-identify"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/binds"
	"github.com/gin-gonic/gin"
)

// Identify plugin init
var Identify = identify.
	New().
	Init(func(iden *identify.Identify) {
		iden.Auth = func(ctx *gin.Context) (interface{}, error) {
			var login binds.Login
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
		iden.FakeTokens = []interface{}{"DEBUG"}
		iden.FakeURLs = []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`}
	})
