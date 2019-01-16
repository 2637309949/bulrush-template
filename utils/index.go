/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package utils

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/utils"
	"path"
	"fmt"
	"os"
)

// GINMODE APP ENV
var GINMODE 	= utils.Some(os.Getenv("GIN_MODE"), "local")
// CONFIGPATH PATH
var CONFIGPATH  = path.Join(".", fmt.Sprintf("conf/%s.yaml", GINMODE))
// Mgo -
var Mgo *bulrush.Mgo
// Rds -
var Rds *bulrush.Rds
// Logger -
var Logger func(string)
// WC -
var WC *bulrush.Config


// Some get or a default value
func Some(target interface{}, initValue interface{}) interface{}{
	if target != nil && target != "" && target != 0 {
		return target
	}
	return initValue
}

func init() {
	WC     = bulrush.NewCfg(CONFIGPATH)
	Logger = bulrush.LoggerWrap(WC)
	Mgo    = bulrush.NewMgo(WC)
	Rds    = bulrush.NewRds(WC)
}