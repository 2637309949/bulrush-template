/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:40
 * @modify date 2019-01-16 20:49:40
 * @desc [description]
 */

package models

import (
	"github.com/2637309949/bulrush-template/addition"
)

// Param info
type Param struct {
	Base  `bson:",inline"`
	Code  string      `bson:"code" form:"code" json:"code" xml:"code"`
	Desc  string      `bson:"desc" form:"desc" json:"desc" xml:"desc"`
	Value interface{} `bson:"value" form:"value" json:"value" xml:"value"`
}

// Register model
func init() {
	addition.Mongo.Register(map[string]interface{}{
		"db":        "test",
		"name":      "param",
		"reflector": &Param{},
	})
}
