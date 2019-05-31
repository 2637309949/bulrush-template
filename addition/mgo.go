/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package addition

import (
	"github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/conf"
)

// Mongo application mongo store
var Mongo = mgo.New(conf.Cfg)
