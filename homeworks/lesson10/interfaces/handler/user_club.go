package handler

import "golang/homeworks/lesson10/application"

type UserClub struct {
	UserClubApp application.UserClubApp
}

func NewUserClub(userClubApp application.UserClubApp) *UserClub {
	return &UserClub{
		UserClubApp: userClubApp,
	}
}