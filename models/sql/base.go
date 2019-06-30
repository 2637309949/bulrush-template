// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gormext"
)

// Base common fields
type Base struct {
	gormext.Model
	Creator    *User `form:"creator" json:"creator" xml:"creator"`
	CreatorID  uint  `form:"_creator" json:"_creator" xml:"_creator"`
	Modifier   *User `form:"modifier" json:"modifier" xml:"modifier"`
	ModifierID uint  `form:"_modifier" json:"_modifier" xml:"_modifier"`
}
