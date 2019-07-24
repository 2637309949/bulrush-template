// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"time"
)

// Now defined Now
func Now() *time.Time {
	now := time.Now()
	return &now
}
