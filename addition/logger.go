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

	"github.com/2637309949/bulrush-addition/logger"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/2637309949/bulrush-template/utils"
)

// Logger application logger
var Logger = logger.CreateLogger(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)))
