// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/gin-gonic/gin"
)

var gormConf = &gormext.Config{}
var _ = conf.Cfg.Unmarshal("sql", gormConf)

// GORMExt application mongo store
var GORMExt = gormext.New(gormConf)

// GORMExt init config
var _ = GORMExt.Init(func(ext *gormext.GORM) {
	ext.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	ext.DB.LogMode(true)
	ext.API.Opts.Prefix = "/template/gorm"
	ext.API.Opts.RouteHooks = &gormext.RouteHooks{
		List: &gormext.ListHook{
			Pre: func(c *gin.Context) {
				Logger.Info("all gormext before")
			},
		},
	}
})
