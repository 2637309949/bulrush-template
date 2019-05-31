/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import (
	"errors"

	identify "github.com/2637309949/bulrush-identify"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/binds"
	"github.com/gin-gonic/gin"
)

// Identify plugin init
var Identify = &identify.Identify{
	Auth: func(ctx *gin.Context) (interface{}, error) {
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
	},
	Model: &identify.RedisModel{
		Redis: addition.Redis,
	},
	FakeTokens: []interface{}{"DEBUG"},
	FakeURLs:   []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
}
