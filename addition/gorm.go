// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"github.com/2637309949/bulrush-addition/gormext"
	"github.com/2637309949/bulrush-template/conf"
)

var gormConf = &gormext.Config{}
var _ = conf.Cfg.Unmarshal("sql", gormConf)

// GORMExt application mongo store
var GORMExt = gormext.New(gormConf)
var _ = GORMExt.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
