/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:21
 * @modify date 2019-01-16 20:50:21
 * @desc [description]
 */

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
