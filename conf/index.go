/**
 * @author [Double]
 * @email [2637309949@qq.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

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

var (
	dir     = "conf/yaml"
	eFlag   = flag.String("e", "", "specify environment variables")
	envFlag = flag.String("env", "", "specify environment variables")
	// ENV APP ENV
	ENV = utils.Some(utils.Some(os.Getenv("ENV"), os.Getenv("env")), "local")
	// CfgPath PATH
	CfgPath string
	// Cfg app cgf
	Cfg *bulrush.Config
)

func init() {
	flag.Parse()
	ENV = utils.Some(utils.Some(*eFlag, *envFlag), ENV)
	envFileName := fmt.Sprintf("%s.yaml", ENV)
	files, err := filepath.Glob(dir + "/**.yaml")
	if err != nil {
		panic(err)
	}
	envFiles := funk.Map(files, func(file string) string {
		base := filepath.Base(file)
		return base
	}).([]string)

	findFrom := funk.Find(envFiles, func(file string) bool {
		return file == envFileName
	})

	if findFrom == nil {
		panic(fmt.Errorf("envFileName %s from env or flag has no been found", envFileName))
	} else {
		CfgPath = path.Join(".", dir, envFileName)
		Cfg = bulrush.NewCfg(CfgPath)
	}
}
