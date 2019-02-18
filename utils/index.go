/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	addition "github.com/2637309949/bulrush-addition"
)

// GINMODE APP ENV
var GINMODE = Some(os.Getenv("GIN_MODE"), "local")

// CONFIGPATH PATH
var CONFIGPATH = path.Join(".", fmt.Sprintf("conf/%s.yaml", GINMODE))

// Mgo -
var Mgo *addition.Mgo

// Rds -
var Rds *addition.Rds

// WC -
var WC *bulrush.Config

// Some get or a default value
func Some(target interface{}, initValue interface{}) interface{} {
	if target != nil && target != "" && target != 0 {
		return target
	}
	return initValue
}

func init() {
	WC = bulrush.NewCfg(CONFIGPATH)
	Mgo = addition.NewMgo(WC)
	Rds = addition.NewRds(WC)
}
