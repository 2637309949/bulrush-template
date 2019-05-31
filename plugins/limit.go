/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import (
	limit "github.com/2637309949/bulrush-limit"
	"github.com/2637309949/bulrush-template/addition"
)

// Limit Plugin init
var Limit = &limit.Limit{
	Frequency: &limit.Frequency{
		Passages: []string{},
		Rules: []limit.Rule{
			limit.Rule{
				Methods: []string{"GET"},
				Match:   "/api/v1/user*",
				Rate:    1,
			},
			limit.Rule{
				Methods: []string{"GET"},
				Match:   "/api/v1/role*",
				Rate:    2,
			},
		},
		Model: &limit.RedisModel{
			Redis: addition.Redis,
		},
	},
}
