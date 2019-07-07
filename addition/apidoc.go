// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"path"

	"github.com/2637309949/bulrush-addition/apidoc"
)

// APIDoc defined http rest api
var APIDoc = apidoc.New()

// Init after mgo and gorm
func init() {
	APIDoc.
		Doc(path.Join("", "doc")).
		Init(func(api *apidoc.APIDoc) {
			api.Prefix = "/docs"
			api.GORMExt = GORMExt
			api.MGOExt = MGOExt
		})
}
