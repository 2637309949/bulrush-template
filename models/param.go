// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"github.com/2637309949/bulrush-template/addition"
)

// Param info
type Param struct {
	Base  `bson:",inline"`
	Code  string      `bson:"code" form:"code" json:"code" xml:"code"`
	Desc  string      `bson:"desc" form:"desc" json:"desc" xml:"desc"`
	Value interface{} `bson:"value" form:"value" json:"value" xml:"value"`
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "param",
		"reflector": &Param{},
	})
}
