// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"
)

// File defined struct
type File struct {
	Model
	UID    string      `gorm:"comment:'标识'"`
	Name   string      `gorm:"comment:'名称'"`
	Type   string      `gorm:"comment:'类型'"`
	Status uint        `gorm:"comment:'状态'"`
	Size   uint        `gorm:"comment:'大小'"`
	URL    string      `gorm:"comment:'URL'"`
	Path   null.String `gorm:"comment:'路径'"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	Name:      "File",
	Reflector: &File{},
	Opts: &gormext.Opts{
		RouteHooks: &gormext.RouteHooks{
			List: &gormext.ListHook{
				Pre: func(c *gin.Context) {
					addition.Logger.Info("File model pre hook")
				},
			},
		},
	},
})
