/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:40
 * @modify date 2019-01-16 20:49:40
 * @desc [description]
 */

package models

import "github.com/globalsign/mgo/bson"

// Base common fields
type Base struct {
	Creator  bson.ObjectId `bson:"_creator" form:"_creator" json:"_creator" xml:"_creator"`
	Modifier bson.ObjectId `bson:"_modifier" form:"_modifier" json:"_modifier" xml:"_modifier" `
	Created  int           `bson:"_created" form:"_created" json:"_created" xml:"_created"`
	Modified int           `bson:"_modified" form:"_modified" json:"_modified" xml:"_modified"`
	Dr       int           `bson:"_dr" form:"_dr" json:"_dr" xml:"_dr"`
}
