/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/utils"
)

// ENV APP ENV
var ENV = utils.Some(os.Getenv("ENV"), "local")

// CfgPath PATH
var CfgPath = path.Join(".", fmt.Sprintf("conf/yaml/%s.yaml", ENV))

// Cfg app cgf
var Cfg = bulrush.NewCfg(CfgPath)
