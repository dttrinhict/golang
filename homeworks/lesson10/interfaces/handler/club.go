package handler

import "golang/homeworks/lesson10/application"

type Club struct {
	ClubApp application.ClubApp
}

func NewClub(clubApp application.ClubApp) *Club {
	return &Club{
		ClubApp: clubApp,
	}
}