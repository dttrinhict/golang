package handler

import (
	"golang/homeworks/lesson12/application"
)

type User struct {
	UserApp application.UserApp
}

func NewUser(userApp application.UserApp) *User {
	return &User{
		UserApp: userApp,
	}
}