// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	addition "github.com/2637309949/bulrush-addition"
)

// I18N defined global I18N
var I18N = addition.NewI18N()
var _ = I18N.Init(func(i18n *addition.I18N) func() {
	return func() {
		i18n.AddLocales(addition.M{
			"zh_CN": addition.M{
				"sys_hello": "你好",
			},
			"en_US": addition.M{
				"sys_hello": "hello",
			},
			"zh_TW": addition.M{
				"sys_hello": "妳好",
			},
		})
	}
})
