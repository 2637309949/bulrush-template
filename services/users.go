package services

import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush_template/models"
)

// FindUsers users
func FindUsers(match map[string]interface{}) []models.User{
	mongo := bulrush.Mongo
	var users []models.User
	User, _ := mongo.Model("user")
	err := User.Find(match).All(&users)
	if err != nil {
		panic(err)
	}
	return users
}

// AddUsers users
func AddUsers(users [] interface{}) {
	var mongo = bulrush.Mongo
	User, _ := mongo.Model("user")
	err := User.Insert(users...)
	if err != nil {
		panic(err)
	}
}