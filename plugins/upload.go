// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	upload "github.com/2637309949/bulrush-upload"
)

func initUpload() *upload.Upload {
	return upload.New()
}

// Upload Plugin init
var Upload = initUpload()
