/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package addition

import (
	"path"

	"github.com/2637309949/bulrush-template/utils"

	"github.com/2637309949/bulrush-addition/logger"
	"github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-addition/redis"
	"github.com/2637309949/bulrush-template/conf"
)

// Mongo application mongo store
var Mongo = mgo.New(conf.Cfg)

// Redis application redis store
var Redis = redis.New(conf.Cfg)

// Logger application logger
var Logger = logger.CreateLogger(path.Join(".", utils.Some(utils.LeftV(conf.Cfg.String("logs")), "logs").(string)))
