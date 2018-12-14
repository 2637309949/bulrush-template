package services

import (
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush_template/utils"
)

// FindUsers users
func FindUsers(match map[string]interface{}) []models.User{
	var users []models.User
	User, _ := utils.Mgo.Model("user")
	err := User.Find(match).All(&users)
	if err != nil {
		panic(err)
	}
	return users
}

// AddUsers users
func AddUsers(users [] interface{}) {
	User, _ := utils.Mgo.Model("user")
	err := User.Insert(users...)
	if err != nil {
		panic(err)
	}
}