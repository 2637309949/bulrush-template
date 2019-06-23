// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	captcha "github.com/2637309949/bulrush-captcha"
)

func initCaptcha() *captcha.Captcha {
	var c = captcha.New()
	c.Secret = "7658388"
	c.URLPrefix = "/captcha"
	return c
}

// Captcha plugin init
var Captcha = initCaptcha()
