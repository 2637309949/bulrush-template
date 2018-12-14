package utils

import (
	"errors"
	"github.com/2637309949/bulrush_template/binds"
	"github.com/gin-gonic/gin"
)

// Auth -
func Auth(c *gin.Context) (interface{}, error) {
	var login binds.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		return nil, err
	}
	if login.Password == "xx" && login.UserName == "xx" {
		return  map[string] interface{} {
					"id":			"3e4r56u80a55",
					"username": 	login.UserName,
				}, nil
	}
	return nil, errors.New("user authentication failed")
}
