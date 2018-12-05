package utils

import (
	"time"
	"encoding/json"
	"github.com/2637309949/bulrush"
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


// Save -
func Save(token map[string]interface{}) {
	accessToken, _  := token["accessToken"]
	refreshToken, _ := token["refreshToken"]
	value, _ := json.Marshal(token)
	bulrush.Redis.Client.Set("TOKEN:" + accessToken.(string),  value, 2*24*time.Hour)
	bulrush.Redis.Client.Set("TOKEN:" + refreshToken.(string), value, 5*24*time.Hour)
}

// Revoke -
func Revoke(accessToken string) bool {
	status, err := bulrush.Redis.Client.Del("TOKEN:" + accessToken).Result()
	if err != nil {
		return false
	} else if status != 1 {
		return false
	}
	return true
}

// Find -
func Find(accessToken string, refreshToken string) map[string]interface{} {
	var imapGet map[string]interface{}
	var token string
	if accessToken != "" {
		token = accessToken
	} else if refreshToken != "" {
		token = refreshToken
	}
	value, err := bulrush.Redis.Client.Get("TOKEN:" + token).Result()
	json.Unmarshal([]byte(value), &imapGet)
	if err != nil {
		return nil
	}
	return imapGet
}
