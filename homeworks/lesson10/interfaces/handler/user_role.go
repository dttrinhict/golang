package handler

import "golang/homeworks/lesson10/application"

type UserRole struct {
	UserRoleApp application.UserRoleApp
}

func NewUserClub(userRoleApp application.UserRoleApp) *UserRole {
	return &UserRole{
		UserRoleApp: userRoleApp,
	}
}