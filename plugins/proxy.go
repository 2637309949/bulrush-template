/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import proxy "github.com/2637309949/bulrush-proxy"

// Proxy Plugin init
var Proxy = &proxy.Proxy{
	Host:  "https://xxx.com",
	Match: "^/api/v1/proxyTest",
	Map: func(reqPath string) string {
		return reqPath
	},
}
