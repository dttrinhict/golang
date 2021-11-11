package handler

import (
	"golang/homeworks/lesson12/application"
)

type Member struct {
	memberApp application.MemberApp
}

func NewMember(memberApp application.MemberApp) *Member {
	return &Member{
		memberApp: memberApp,
	}
}