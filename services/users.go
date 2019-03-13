/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:39
 * @modify date 2019-01-16 20:50:39
 * @desc [description]
 */

package services

import (
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/models"
)

// FindUsers users
func FindUsers(match map[string]interface{}) []models.User {
	var users []models.User
	User, _ := addition.Mongo.Model("user")
	err := User.Find(match).All(&users)
	if err != nil {
		panic(err)
	}
	return users
}

// AddUsers users
func AddUsers(users []interface{}) {
	User, _ := addition.Mongo.Model("user")
	err := User.Insert(users...)
	if err != nil {
		panic(err)
	}
}
