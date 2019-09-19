// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/conf"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func app() bulrush.Bulrush {
	return addPlugin(bulrush.Default().Config(conf.CPath))
}
