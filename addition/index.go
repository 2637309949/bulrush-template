/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package addition

import (
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-addition/mongo"
	"github.com/2637309949/bulrush-addition/redis"
	"github.com/2637309949/bulrush-template/utils"
)

// GINMODE APP ENV
var GINMODE = utils.Some(os.Getenv("GIN_MODE"), "local")

// CONFIGPATH PATH
var CONFIGPATH = path.Join(".", fmt.Sprintf("conf/%s.yaml", GINMODE))

// Mongo -
var Mongo *mongo.Mongo

// Redis -
var Redis *redis.Redis

// WC -
var WC *bulrush.Config

func init() {
	WC = bulrush.NewCfg(CONFIGPATH)
	Mongo = mongo.New(WC)
	Redis = redis.New(WC)
}
