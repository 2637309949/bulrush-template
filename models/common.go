// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import "github.com/globalsign/mgo/bson"

// Base common fields
type Base struct {
	Creator  bson.ObjectId `bson:"_creator" form:"_creator" json:"_creator" xml:"_creator"`
	Modifier bson.ObjectId `bson:"_modifier" form:"_modifier" json:"_modifier" xml:"_modifier" `
	Created  int           `bson:"_created" form:"_created" json:"_created" xml:"_created"`
	Modified int           `bson:"_modified" form:"_modified" json:"_modified" xml:"_modified"`
	Dr       int           `bson:"_dr" form:"_dr" json:"_dr" xml:"_dr"`
}
