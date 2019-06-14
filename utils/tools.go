// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

// Some get or a default value
func Some(target interface{}, initValue interface{}) interface{} {
	if target != nil && target != "" && target != 0 {
		return target
	}
	return initValue
}

// LeftV get left value
func LeftV(left interface{}, right interface{}) interface{} {
	return left
}

// RightV get right value
func RightV(left interface{}, right interface{}) interface{} {
	return right
}
