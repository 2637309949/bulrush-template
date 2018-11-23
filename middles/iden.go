package middles

import (
	"errors"
	"github.com/2637309949/bulrush/middles"
	"github.com/2637309949/bulrush_template/binds"
	"github.com/gin-gonic/gin"
)

var store = make([]map[string]interface{}, 0)

// save token
func save(token map[string]interface{}) {
	store = append(store, token)
}

// revoke token
func revoke(accessToken string) bool {
	var index = -1
	for i, item := range store {
		if value, _ := item["accessToken"]; value == accessToken {
			index = i
			break
		}
	}
	if index != -1 {
		store = append(store[:index], store[index+1:]...)
		return true
	}
	return false
}

// find token by accessToken or refleshToken
func find(accessToken string, refreshToken string) map[string]interface{} {
	var target map[string]interface{}
	for _, item := range store {
		accessTokenMatch := func(item map[string]interface{}, accessToken string) bool{
			value, _ := item["accessToken"]
			match := value == accessToken
			return match
		}
		refreshTokenMatch := func(item map[string]interface{}, refreshToken string) bool{
			value, _ := item["refreshToken"]
			match := value == refreshToken
			return match
		}
		allMatch := func(item map[string]interface{}) bool{
			if accessToken != "" && refreshToken != "" {
				return accessTokenMatch(item, accessToken) && refreshTokenMatch(item, refreshToken)
			} else if accessToken != "" {
				return accessTokenMatch(item, accessToken)
			} else if refreshToken != "" {
				return refreshTokenMatch(item, refreshToken)
			}
			return false
		}
		if match := allMatch(item); match {
			target = item
			break
		}
	}
	return target
}


// idenEntity for auth
var entity = &middles.Iden {
	ExpiresIn: 86400,
	Routes: middles.RoutesGroup {
		ObtainTokenRoute:  "/obtainToken",
		RevokeTokenRoute:  "/revokeToken",
		RefleshTokenRoute: "/refleshToken",
	},
	Auth: func(c *gin.Context) (interface{}, error) {
		var login binds.Login
		if err := c.ShouldBindJSON(&login); err != nil {
			return nil, err
		}
		if login.Password == "xx" && login.UserName == "xx" {
			return map[string] interface{} {
				"id":			"3e4r56u80a55",
				"username": 	login.UserName,
			}, nil
		}
		return nil, errors.New("user authentication failed")
	},
	Tokens: middles.TokensGroup {
		Save 	:save,
		Revoke  :revoke,
		Find	:find,
	},
	IgnoreURLs: []interface{}{ `^/api/v1/ignore$`, `^/api/v1/docs/*` },
}

// iden as Function that accept some injection parameters
func iden (injects map[string]interface{}) {
	entity.Inject(injects)
}