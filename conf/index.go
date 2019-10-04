// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/utils"
)

// var declare
var (
	dir     = "conf/yaml"
	eFlag   = flag.String("e", "", "specify environment variables")
	envFlag = flag.String("env", "", "specify environment variables")
	ENV     = utils.Some(utils.Some(os.Getenv("ENV"), os.Getenv("env")), "local")
	CPath   string
	Conf    *bulrush.Config
)

func init() {
	ENV = utils.Some(utils.Some(*eFlag, *envFlag), ENV)
	CPath = path.Join(".", dir, fmt.Sprintf("%s.yaml", ENV))
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		panic(fmt.Errorf("file %s not existed", CPath))
	}
	Conf = bulrush.LoadConfig(CPath)
}
