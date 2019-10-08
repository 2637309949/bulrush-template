// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/utils"
)

// ENV defined env var
var ENV = utils.Some(utils.Some(utils.Some(os.Getenv("ENV"), os.Getenv("env")), os.Getenv("e")), "local")

// CPath defined config path
var CPath = path.Join(".", "conf/yaml", fmt.Sprintf("%s.yaml", ENV))

// Conf defined config
var Conf = bulrush.LoadConfig(CPath)
