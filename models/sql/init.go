// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"github.com/2637309949/bulrush-template/addition"
)

func init() {
	addition.GORMExt.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	// 1. create user
	tx := addition.GORMExt.DB.Begin()
	tx.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.Exec("DROP DATABASE IF EXISTS TEST;")
	tx.Exec("CREATE DATABASE TEST;")
	tx.DropTable(&User{})
	tx.CreateTable(&User{})
	first := &User{
		Name: "L1212",
		Age:  23,
	}
	tx.Create(first)
	second := &User{
		Name: "L1212",
		Age:  24,
		Base: Base{
			CreatorID: first.ID,
		},
	}
	tx.Create(second)

	// 2. create permission
	tx.DropTable(&Permission{})
	tx.CreateTable(&Permission{})
	tx.Create(&Permission{
		Code: "ER5T12",
		Name: "FINANCE MENU",
		Type: 2,
	})
	tx.Create(&Permission{
		Code: "ER5T15",
		Name: "MENU",
		Type: 1,
		Pid:  1,
	})

	// 3. create role
	tx.DropTable(&Role{})
	tx.CreateTable(&Role{})
	tx.Create(&Role{
		Name: "FINANCE1",
		Type: 1,
	})

	p1 := &Permission{}
	p1.ID = 1
	tx.Create(&Role{
		Name:        "FINANCE1",
		Type:        1,
		Permissions: []*Permission{p1},
	})

	tx.Exec("SET FOREIGN_KEY_CHECKS = 1;")
	err := tx.Commit().Error
	if err != nil {
		err = tx.Rollback().Error
	}
}
