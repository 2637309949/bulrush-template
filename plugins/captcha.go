// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	captcha "github.com/2637309949/bulrush-captcha"
	"github.com/mojocn/base64Captcha"
)

// Captcha plugin init
var Captcha = &captcha.Captcha{
	URLPrefix: "/captcha",
	Secret:    "7658388",
	Config: base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	},
}
