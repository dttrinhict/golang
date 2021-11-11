package handler

import "golang/homeworks/lesson12/application"

type MemberClub struct {
	MemberClubApp application.MemberClubApp
}

func NewMemberClub(memberClubApp application.MemberClubApp) *MemberClub {
	return &MemberClub{
		MemberClubApp: memberClubApp,
	}
}