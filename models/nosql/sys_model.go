// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/globalsign/mgo/bson"
)

// Model common fields
type Model struct {
	mgoext.Model `bson:",inline"`
	CreatorID    bson.ObjectId `bson:"_creator,omitempty" br:"comment:'创建人ID'"`
	Creator      *User         `bson:"creator,omitempty" br:"ref(user,_creator,_id),comment:'创建人',up(password)"`
	ModifierID   bson.ObjectId `bson:"_modifier,omitempty" br:"comment:'修改人ID'"`
	Modifier     *User         `bson:"modifier,omitempty" br:"ref(user,_modifier,_id),comment:'修改人',up(password)"`
	DeleterID    bson.ObjectId `bson:"_deleter,omitempty" br:"comment:'删除人ID'"`
	Deleter      *User         `bson:"deleter,omitempty"  br:"ref(user,_deleter,_id),comment:'删除人',up(password)"`
}
