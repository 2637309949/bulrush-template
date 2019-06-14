// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
