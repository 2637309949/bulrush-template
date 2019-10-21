// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	captcha "github.com/2637309949/bulrush-captcha"
)

// Captcha plugin init
var Captcha = captcha.
	New().
	Init(func(c *captcha.Captcha) {
		c.Secret = "7658388"
		c.URLPrefix = "/captcha"
	})
