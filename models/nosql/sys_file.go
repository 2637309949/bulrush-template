// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"
)

// File defined struct
type File struct {
	Model  `bson:",inline"`
	UID    string      `bson:"uid" br:"comment:'标识'"`
	Name   string      `bson:"name" br:"comment:'名称'"`
	Type   string      `bson:"type" br:"comment:'类型'"`
	Status uint        `bson:"status" br:"comment:'状态'"`
	Size   uint        `bson:"size" br:"comment:'大小'"`
	URL    string      `bson:"url" br:"comment:'URL'"`
	Path   null.String `bson:"path" br:"comment:'路径'"`
}

var _ = MGOExt.Register(&mgoext.Profile{
	Name:      "File",
	Reflector: new(File),
	Opts: &mgoext.Opts{
		RouteHooks: &mgoext.RouteHooks{
			List: &mgoext.ListHook{
				Pre: func(c *gin.Context) {
					Logger.Info("file before")
				},
			},
		},
	},
})
