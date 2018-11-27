package middles

import (
	"time"
	"encoding/json"
	"errors"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/middles"
	"github.com/2637309949/bulrush_template/binds"
	"github.com/gin-gonic/gin"
)

var store = make([]map[string]interface{}, 0)

// save token
func save(token map[string]interface{}) {
	accessToken, _  := token["accessToken"]
	refreshToken, _ := token["refreshToken"]
	value, _ := json.Marshal(token)
	bulrush.Redis.Client.Set("TOKEN:" + accessToken.(string),  value, 2*24*time.Hour)
	bulrush.Redis.Client.Set("TOKEN:" + refreshToken.(string), value, 5*24*time.Hour)
}

// revoke token
func revoke(accessToken string) bool {
	status, err := bulrush.Redis.Client.Del("TOKEN:" + accessToken).Result()
	if err != nil {
		return false
	} else if status != 1 {
		return false
	}
	return true
}

// find token by accessToken or refleshToken
func find(accessToken string, refreshToken string) map[string]interface{} {
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

var identity = &middles.Iden {
	ExpiresIn: 	86400,
	Routes: 	middles.RoutesGroup {
					ObtainTokenRoute:  "/obtainToken",
					RevokeTokenRoute:  "/revokeToken",
					RefleshTokenRoute: "/refleshToken",
				},
	Auth: 		func(c *gin.Context) (interface{}, error) {
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
				},
	Tokens: 	middles.TokensGroup {
					Save 	:save,
					Revoke  :revoke,
					Find	:find,
				},
	IgnoreURLs: []interface{}{ `^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$` },
}
