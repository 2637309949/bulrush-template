package main

import (
	"github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/utils"
	"github.com/2637309949/bulrush_template/routes"
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush_template/middles"
	"path"
	"time"
	"fmt"
	"os"
)

// GINMODE APP ENV
var GINMODE = utils.Some(os.Getenv("GIN_MODE"), "local")
// CONFIGPATH PATH
var CONFIGPATH = path.Join(fmt.Sprintf("conf/%s.yaml", GINMODE))

func main() {
	app := bulrush.New()
	app.Use(func(c *gin.Context) {
		out 	 := os.Stdout
		start 	 := time.Now()
		path 	 := c.Request.URL.Path
		raw 	 := c.Request.URL.RawQuery
		c.Next()
		end 	 := time.Now()
		latency  := end.Sub(start)
		clientIP := c.ClientIP()
		method   := c.Request.Method
		if raw != "" {
			path = path + "?" + raw
		}
		fmt.Fprintf(out, "--> %v %8v %s %6s %s\n", end.Format("2006/01/02 15:04:05"), latency, clientIP, method, path)
	})
	app.LoadConfig(CONFIGPATH, bulrush.YAMLMode)
	app.Inject(&middles.Middles{}, &routes.Routes{}, &models.Model{})
	app.Run()
}
