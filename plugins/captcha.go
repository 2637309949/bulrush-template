/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

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
