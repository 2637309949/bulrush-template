// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/2637309949/bulrush-template/addition"
)

// FindUsers users
// users := []User{} 需要具体化时用Type去声明变量，否则直接用反射体
func FindUsers(match map[string]interface{}) (interface{}, error) {
	User := addition.MGOExt.Model("User")
	users := addition.MGOExt.Vars("User")
	if err := User.Find(match).All(users); err != nil {
		return nil, err
	}
	return users, nil
}

// AddUsers users
func AddUsers(users []interface{}) {
	User := addition.MGOExt.Model("User")
	User.Insert(users...)
}
