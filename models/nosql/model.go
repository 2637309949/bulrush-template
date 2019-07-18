// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	"time"

	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/globalsign/mgo/bson"
)

// Model common fields
type Model struct {
	mgoext.Model `bson:",inline"`
	CreatorID    bson.ObjectId `bson:"_creator,omitempty" br:"comment:'创建人ID'"`
	Creator      *User         `bson:"creator,omitempty" br:"ref(user,_creator,_id),comment:'创建人',up(password)"`
	UpdatorID    bson.ObjectId `bson:"_updator,omitempty" br:"comment:'修改人ID'"`
	Updator      *User         `bson:"modifier,omitempty" br:"ref(user,_updator,_id),comment:'修改人',up(password)"`
	DeleterID    bson.ObjectId `bson:"_deleter,omitempty" br:"comment:'删除人ID'"`
	Deleter      *User         `bson:"deleter,omitempty"  br:"ref(user,_deleter,_id),comment:'删除人',up(password)"`
}

// PresetModel defined Preset User
// 系统内置数据时的默认参数
func PresetModel() Model {
	now := time.Now()
	return Model{
		Model: mgoext.Model{
			UpdatedAt: &now,
			CreatedAt: &now,
		},
		CreatorID: bson.ObjectIdHex("5d2fdc047dead1c7924b3a52"),
		UpdatorID: bson.ObjectIdHex("5d2fdc047dead1c7924b3a52"),
	}
}
