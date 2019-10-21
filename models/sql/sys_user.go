// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/2637309949/bulrush"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/utils"
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/guregu/null.v3"
)

// User defined struct
type User struct {
	Model
	Name     string                 `gorm:"comment:'名称';unique;not null"`
	Password string                 `gorm:"comment:'密码';not null"`
	Salt     string                 `gorm:"comment:'盐噪点';not null"`
	Age      uint                   `gorm:"comment:'年龄'"`
	Birthday *time.Time             `gorm:"comment:'生日'"`
	Mobile   string                 `gorm:"comment:'手机'"`
	Email    null.String            `gorm:"comment:'邮箱'"`
	Roles    []*Role                `gorm:"comment:'角色列表';many2many:role2users;"`
	Ignored  *struct{ Name string } `gorm:"-"`
}

// SetPassword Method to set salt and hash the password for a user
func (u *User) SetPassword(password string) {
	b, err := utils.RandomBytes(16)
	if err != nil {
		panic(err)
	}
	u.Salt = fmt.Sprintf("%x", b)
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt), 1000, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	u.Password = fmt.Sprintf("%x", dk)
}

// ValidPassword Method to check the entered password is correct or not
func (u *User) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt), 1000, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return u.Password == fmt.Sprintf("%x", dk)
}

// Exist defined user is existed or not
func (u *User) Exist(query interface{}, args ...interface{}) bool {
	count := 0
	GORMExt.Model("User").Where(query, args...).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

var _ = GORMExt.Register(&gormext.Profile{
	Name:      "User",
	Reflector: &User{},
	Opts: &gormext.Opts{
		RouteHooks: &gormext.RouteHooks{
			List: &gormext.ListHook{
				Pre: func(c *gin.Context) {
					Logger.Info("User model pre hook")
				},
			},
		},
	},
})

/**
 * @api {post} /signup                        用户注册
 * @apiGroup Users
 * @apiDescription                            用户注册
 * @apiSuccess        {Object}  Mess		  实体类
 * @apiSuccess        {String}  Mess.message  消息内容
 * @apiSuccessExample {json}                  正常返回
 * HTTP/1.1 200 OK
 * {
 *    "message": "ok"
 * }
 * @apiErrorExample {json}                    错误返回
 * HTTP/1.1 500 Internal ServerError
 * {
 *    "message": "the username already exists"
 * }
**/
func signup(r *gin.RouterGroup) {
	r.POST("/signup", func(c *gin.Context) {
		ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
			var form = &User{}
			if err := c.ShouldBindBodyWith(form, binding.JSON); err != nil {
				Logger.Error(err.Error())
				return nil, err
			}
			form.SetPassword(form.Password)
			if form.Exist("name = ?", form.Name) {
				return nil, errors.New("the username already exists")
			}
			if err := GORMExt.DB.Save(form).Error; err != nil {
				return nil, err
			}
			return gin.H{
				"message": "ok",
			}, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, ret)
	})
}

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup, ri *bulrush.ReverseInject) {
	ri.Register(signup)

	GORMExt.API.List(r, "User").Post(func(c *gin.Context) {
		Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	}).RouteHooks(&gormext.RouteHooks{
		// Override global config, never query only by own
		List: &gormext.ListHook{
			Cond: func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
				return cond
			},
		},
	})
	// Example Restful Api
	GORMExt.API.Feature("subUser").List(r, "User")
	GORMExt.API.One(r, "User")
	GORMExt.API.Create(r, "User")
	GORMExt.API.Update(r, "User")
	GORMExt.API.Delete(r, "User")
}
