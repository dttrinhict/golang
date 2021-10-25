package application

import (
	domainservice "golang/homeworks/lesson10/domain/services"
)

type UserClubImpl struct {
	domainUserClubService domainservice.DomainUserClubService
}

type UserClubApp interface {
	AssignUserToClub(club Club) (users []User, err error)
	AssignClubsToUser(user User) (clubs []Club, err error)
	GetUsersOfClub(club Club) (users []User, err error)
	GetClubsOfUser(user User) (clubs []Club, err error)
	RemoveUserFromClub(club Club) (users []User, err error)
}

func User_Club_App(domainUserClubService domainservice.DomainUserClubService) UserClubApp {
	return &UserClubImpl{
		domainUserClubService: domainUserClubService,
	}
}

func (u UserClubImpl) AssignUserToClub(club Club) (users []User, err error) {
	domainClub := MapClubAppToDomain(club)
	domainUsers, err := u.domainUserClubService.AssignUserToClub(domainClub)
	if err != nil {
		return users, err
	}
	return MapUsersApp(domainUsers), err
}

func (u UserClubImpl) AssignClubsToUser(user User) (clubs []Club, err error) {
	panic("implement me")
}

func (u UserClubImpl) GetUsersOfClub(club Club) (users []User, err error) {
	panic("implement me")
}

func (u UserClubImpl) GetClubsOfUser(user User) (clubs []Club, err error) {
	panic("implement me")
}

func (u UserClubImpl) RemoveUserFromClub(club Club) (users []User, err error) {
	panic("implement me")
}