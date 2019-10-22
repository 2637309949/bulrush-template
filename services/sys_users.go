// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

// FindUsers users
// users := []User{} 需要具体化时用Type去声明变量，否则直接用反射体
func (srv *Services) FindUsers(match map[string]interface{}) (interface{}, error) {
	User := srv.MGOExt.Model("User")
	users := srv.MGOExt.Vars("User")
	if err := User.Find(match).All(users); err != nil {
		return nil, err
	}
	return users, nil
}

// AddUsers users
func (srv *Services) AddUsers(users []interface{}) {
	User := srv.MGOExt.Model("User")
	User.Insert(users...)
}
