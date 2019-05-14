/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"fmt"
	"strings"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bulrush.Default()
	app.Config(conf.CfgPath)
	app.Inject("bulrushApp")
	appUsePlugins(app)
	app.Run(func(httpProxy *gin.Engine, config *bulrush.Config) {
		port := config.GetString("port", ":8080")
		port = strings.TrimSpace(port)
		name := config.GetString("name", "")
		if prefix := port[:1]; prefix != ":" {
			port = fmt.Sprintf(":%s", port)
		}
		fmt.Println("\n\n================================")
		fmt.Printf("App: %s\n", name)
		fmt.Printf("Listen on %s\n", port)
		fmt.Println("================================")
		httpProxy.Run(port)
	})
}
