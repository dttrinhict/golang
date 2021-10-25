package handler

import (
	"golang/homeworks/lesson10/application"
)

type User struct {
	UserApp application.UserApp
}

func NewUser(userApp application.UserApp) *User {
	return &User{
		UserApp: userApp,
	}
}