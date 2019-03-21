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

// Mongo application mongo store
var Mongo *mongo.Mongo

// Redis application redis store
var Redis *redis.Redis

// Config application cfg
var Config *bulrush.Config

func init() {
	Config = bulrush.NewCfg(path.Join(".", fmt.Sprintf("conf/%s.yaml", utils.Some(os.Getenv("GIN_MODE"), "local"))))
	Mongo = mongo.New(Config)
	Redis = redis.New(Config)
}
