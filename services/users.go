// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/2637309949/bulrush-template/addition"
)

// FindUsers users
func FindUsers(match map[string]interface{}) interface{} {
	// users := []User{} 需要具体化时用Type去声明变量，否则直接用反射体
	users := addition.Mongo.Vars("user")
	User := addition.Mongo.Model("user")
	User.Find(match).All(&users)
	return users
}

// AddUsers users
func AddUsers(users []interface{}) {
	User := addition.Mongo.Model("user")
	User.Insert(users...)
}
