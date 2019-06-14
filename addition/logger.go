// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"path"

	"github.com/2637309949/bulrush-addition/logger"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/2637309949/bulrush-template/utils"
)

// Logger application logger
var Logger = logger.CreateLogger(
	logger.INFOLevel,
	nil,
	[]*logger.Transport{
		// only for error
		&logger.Transport{
			Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "error"),
			Level:   logger.ERRORLevel,
			Maxsize: logger.Maxsize,
		},
		// combined all level
		&logger.Transport{
			Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "combined"),
			Level:   logger.INFOLevel,
			Maxsize: logger.Maxsize,
		},
		// console level
		&logger.Transport{
			Level: logger.INFOLevel,
		},
	},
)
