// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
