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

	delivery "github.com/2637309949/bulrush-delivery"
)

// Delivery Upload, Logger, Captcha Plugin init
var Delivery = &delivery.Delivery{
	URLPrefix: "/public",
	Path:      path.Join("assets/public", ""),
}
