// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/utils"
	"github.com/thoas/go-funk"
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
	var (
		files []string
		err   error
	)
	ENV = utils.Some(utils.Some(*eFlag, *envFlag), ENV)
	fileName := fmt.Sprintf("%s.yaml", ENV)
	if files, err = filepath.Glob(dir + "/**.yaml"); err != nil {
		panic(err)
	}
	if funk.Find(files, func(file string) bool {
		return filepath.Base(file) == fileName
	}) == nil {
		panic(fmt.Errorf("envFileName %s from env or flag has no been found", fileName))
	}
	CPath = path.Join(".", dir, fileName)
	Conf = bulrush.LoadConfig(CPath)
}
