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
var WC *bulrush.WellCfg
func init() {
	WC     = bulrush.NewWc(CONFIGPATH)
	Logger = bulrush.LoggerWrap(WC)
	Mgo    = bulrush.NewMgo(WC)
	Rds    = bulrush.NewRds(WC)
}