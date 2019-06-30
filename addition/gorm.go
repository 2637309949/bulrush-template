// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-template/conf"
)

// GORMExt application mongo store
var GORMExt = gormext.New(conf.Cfg)
var _ = GORMExt.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
