package handler

import (
	"golang/homeworks/lesson10/application"
)

type Member struct {
	memberApp application.MemberApp
}

func NewMember(memberApp application.MemberApp) *Member {
	return &Member{
		memberApp: memberApp,
	}
}