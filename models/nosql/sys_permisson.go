// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/globalsign/mgo/bson"
)

// Permission info
type Permission struct {
	Model       `bson:",inline"`
	Code        string        `bson:"code" br:"comment:'编码'"`
	Name        string        `bson:"name" br:"comment:'名称'"`
	Pid         bson.ObjectId `bson:"pid,omitempty" br:"comment:'父级ID'"`
	PPermission *Permission   `bson:"ppermission,omitempty" br:"ref(permission,pid,_id)'"`
	Type        uint          `bson:"type" br:"comment:'类型',enum:'一级菜单=101 二级菜单=102 三级菜单=103 按钮=104 自定义=105'"`
}

var _ = MGOExt.Register(&mgoext.Profile{
	Name:      "Permission",
	Reflector: new(Permission),
}).Init(func(ext *mgoext.Mongo) {
	(&Param{}).
		AddEnum("Permission", "Type", []Value{
			Value{
				Key:   "一级菜单",
				Value: "101",
			},
			Value{
				Key:   "二级菜单",
				Value: "102",
			},
			Value{
				Key:   "三级菜单",
				Value: "103",
			},
			Value{
				Key:   "按钮",
				Value: "104",
			},
			Value{
				Key:   "自定义",
				Value: "105",
			},
		})
})
