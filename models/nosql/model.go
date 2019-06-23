// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	"github.com/2637309949/bulrush-addition/mgo"
	"github.com/globalsign/mgo/bson"
)

// Model common fields
type Model struct {
	mgo.Model `bson:",inline"`
	Creator   bson.ObjectId `ref:"user;up(password,age)" bson:"_creator" form:"_creator" json:"_creator" xml:"_creator"`
	Modifier  bson.ObjectId `ref:"user" bson:"_modifier" form:"_modifier" json:"_modifier" xml:"_modifier" `
}