/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import (
	"path"

	upload "github.com/2637309949/bulrush-upload"
)

// Upload Plugin init
var Upload = &upload.Upload{
	URLPrefix: "/public/upload",
	Path:      path.Join("assets/public/upload", ""),
}
