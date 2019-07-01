// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	mgoext "github.com/2637309949/bulrush-addition/mgoext"
	"github.com/2637309949/bulrush-template/conf"
)

var mgoConf = &mgoext.Config{}
var _ = conf.Cfg.Unmarshal("mongo", mgoConf)

// MGOExt application mongo store
var MGOExt = mgoext.New(mgoConf)
var _ = MGOExt.Init(func(ext *mgoext.Mongo) {
	ext.API.Opts.Prefix = "/template/mgo"
})
