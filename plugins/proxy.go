// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import proxy "github.com/2637309949/bulrush-proxy"

// Proxy Plugin init
var Proxy = &proxy.Proxy{
	Host:  "https://www.baidu.com",
	Match: "^/proxy",
	Map: func(reqPath string) string {
		return reqPath
	},
}
