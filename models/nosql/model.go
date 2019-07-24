// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	"time"

	"github.com/2637309949/bulrush-template/utils"
	"github.com/globalsign/mgo/bson"
)

// Model common fields
type Model struct {
	ID        bson.ObjectId `bson:"_id,omitempty" br:"comment:'模型ID'"`
	CreatorID bson.ObjectId `bson:"_creator,omitempty" br:"comment:'创建人ID'"`
	Creator   *User         `bson:"creator,omitempty" br:"ref(users,_creator,_id),comment:'创建人',up(password)"`
	CreatedAt *time.Time    `bson:"_created" br:"comment:'创建时间'"`

	UpdatorID bson.ObjectId `bson:"_updator,omitempty" br:"comment:'修改人ID'"`
	Updator   *User         `bson:"modifier,omitempty" br:"ref(users,_updator,_id),comment:'修改人',up(password)"`
	UpdatedAt *time.Time    `bson:"_updated" br:"comment:'修改时间'"`

	DeleterID bson.ObjectId `bson:"_deleter,omitempty" br:"comment:'删除人ID'"`
	Deleter   *User         `bson:"deleter,omitempty"  br:"ref(users,_deleter,_id),comment:'删除人',up(password)"`
	DeletedAt *time.Time    `bson:"_deleted" br:"comment:'删除时间'"`
}

// DefaultModel defined DefaultModel
func DefaultModel() Model {
	return Model{CreatedAt: utils.Now(), UpdatedAt: utils.Now()}
}
