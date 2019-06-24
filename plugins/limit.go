// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
				Rate:    15,
			},
			limit.Rule{
				Methods: []string{"GET"},
				Match:   "/api/v1/role*",
				Rate:    5,
			},
		},
		Model: &limit.RedisModel{
			Redis: addition.Redis,
		},
	},
}
