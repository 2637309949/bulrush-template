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

	"github.com/2637309949/bulrush-template/conf"

	"github.com/2637309949/bulrush-addition/mongo"
	"github.com/2637309949/bulrush-addition/redis"
)

// Mongo application mongo store
var Mongo *mongo.Mongo

// Redis application redis store
var Redis *redis.Redis

func init() {
	fmt.Println(conf.Cfg)
	Mongo = mongo.New(conf.Cfg)
	Redis = redis.New(conf.Cfg)
}
