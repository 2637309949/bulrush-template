// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"net/http"
	"fmt"
	"time"
	"golang.org/x/crypto/scrypt"
	"github.com/2637309949/bulrush-utils/funcs"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"
)

// User defined struct
type User struct {
	Model
	Name     string                 `gorm:"comment:'名称';unique;not null"`
	Password string                 `gorm:"comment:'密码';not null"`
	Salt	 string				    `gorm:"comment:'盐噪点';not null"`
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

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "User",
	Reflector: &User{},
	Opts: &gormext.Opts{
		RouteHooks: &gormext.RouteHooks{
			List: &gormext.ListHook{
				Pre: func(c *gin.Context) {
					addition.Logger.Info("User model pre hook")
				},
			},
		},
	},
})

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	addition.GORMExt.API.List(r, "User").Post(func(c *gin.Context) {
		addition.Logger.Info("after")
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
	addition.GORMExt.API.Feature("subUser").List(r, "User")
	addition.GORMExt.API.One(r, "User")
	addition.GORMExt.API.Create(r, "User")
	addition.GORMExt.API.Update(r, "User")
	addition.GORMExt.API.Delete(r, "User")

	// Model Api
	r.POST("/signup", func(c *gin.Context) {
		ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
			var newUser = &User{}
			if err := c.ShouldBindJSON(newUser); err != nil {
				addition.Logger.Error(err.Error())
				return nil, err
			}
			newUser.SetPassword(newUser.Password)
			if err := addition.GORMExt.DB.Save(newUser).Error; err != nil {
				return nil, err
			}
			return gin.H{
				"message": "User added succesfully.",
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
