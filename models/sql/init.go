// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	"github.com/2637309949/bulrush_template/addition"
)

func init() {
	// 1.重新创建User
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
		Age:  23,
		Base: Base{
			Creator: first,
		},
	}
	tx.Create(second)
	tx.Exec("SET FOREIGN_KEY_CHECKS = 1;")
	tx.Commit()

}
