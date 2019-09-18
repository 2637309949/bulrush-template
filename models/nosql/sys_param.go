// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/thoas/go-funk"
)

// Value defined struct
type Value struct {
	Category    string `bson:"category" br:"comment:'类别'"`
	SubCategory string `bson:"subcategory" br:"comment:'子类别'"`
	Key         string `bson:"key" br:"comment:'属性'"`
	Value       string `bson:"value" br:"comment:'属性值'"`
}

// Param defined struct
type Param struct {
	Model `bson:",inline"`
	Code  string  `bson:"code" br:"comment:'编码'"`
	Name  string  `bson:"name" br:"comment:'名称'"`
	Value []Value `bson:"value" br:"comment:'属性值'"`
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	Name:      "Param",
	Reflector: &Param{},
	AutoHook:  true,
}).Init(func(ext *mgoext.Mongo) {
	Model := addition.MGOExt.Model("Param")
	for _, key := range []string{"code", "name"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := Model.EnsureIndex(index); err != nil {
			addition.Logger.Error(err.Error())
		}
	}
})

// AddEnum defined add enum type
func (p *Param) AddEnum(model string, key string, value []Value) *Param {
	Model := addition.MGOExt.Model("Param")
	enum := &Param{Code: "enum"}
	if err := Model.Find(bson.M{"code": "enum"}).One(enum); err == mgo.ErrNotFound {
		enum.Model = DefaultModel()
		if err := Model.Insert(enum); err != nil {
			addition.Logger.Error(err.Error())
			return p
		}
	}
	for _, v := range value {
		v.Category = model
		v.SubCategory = key
		one := funk.Find(enum.Value, func(item Value) bool {
			return item.Category == v.Category && item.SubCategory == v.SubCategory && item.Key == v.Key
		})
		if one == nil {
			enum.Value = append(enum.Value, Value{
				Category:    v.Category,
				SubCategory: v.SubCategory,
				Key:         v.Key,
				Value:       v.Value,
			})
		}
	}
	if err := Model.Update(bson.M{"code": "enum"}, enum); err != nil {
		addition.Logger.Error(err.Error())
	}
	return p
}
