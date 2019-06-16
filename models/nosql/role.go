// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	"github.com/2637309949/bulrush-template/addition"
	"github.com/globalsign/mgo/bson"
)

// Role info
type Role struct {
	Base        `bson:",inline"`
	Name        string          `bson:"name" form:"name" json:"name" xml:"name"`
	Type        string          `bson:"type" form:"type" json:"type" xml:"type"`
	Permissions []bson.ObjectId `bson:"permissions" form:"permissions" json:"permissions" xml:"permissions" `
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "role",
		"reflector": &Role{},
	})
}
