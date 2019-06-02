/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/conf"
)

// InitApp init applications
func InitApp() bulrush.Bulrush {
	app := bulrush.Default()
	app.Config(conf.CfgPath)
	app.Inject("bulrushApp")
	appUsePlugins(app)
	return app
}
