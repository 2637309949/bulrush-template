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

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/conf"
)

func main() {
	app := bulrush.Default()
	app.Config(conf.CfgPath)
	app.Inject("bulrushApp")
	appUsePlugins(app)
	app.Run(func(err error, config *bulrush.Config) {
		if err != nil {
			panic(err)
		} else {
			name := config.GetString("name", "")
			port := config.GetString("port", "")
			fmt.Println("\n\n================================")
			fmt.Printf("App: %s\n", name)
			fmt.Printf("Listen on %s\n", port)
			fmt.Println("================================")
		}
	})
}
