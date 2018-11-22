package routes

import (
	"net/http"
	"github.com/globalsign/mgo/bson"
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush_template/services"
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush"
)

var db = make(map[string]string)
// Hello as Function that accept some injection parameters
func Hello (engine *gin.Engine, r *gin.RouterGroup, mongo *bulrush.MongoGroup) {
	r.GET("/ping", func(c *gin.Context) {
		services.AddUsers([]interface {} {
			models.User{
				Name: "double",
				Password: "111111",
				Age: 24,
			},
		})
		users := services.FindUsers(bson.M{"name": "double"})
		c.JSON(http.StatusOK, gin.H{
			"data": 	users,
			"errcode": 	nil,
			"errmsg": 	nil,
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": 	"ok",
			"errcode": 	nil,
			"errmsg": 	nil,
		})
	})
}