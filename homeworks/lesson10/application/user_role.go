package application

import (
	domainservice "golang/homeworks/lesson10/domain/services"
)

type UserRoleImpl struct {
	domainUserRoleService domainservice.DomainUserRoleService
}

type UserRoleApp interface {
	AssignUserToRole(role Role) (users []User, err error)
	AssignClubsToUser(user User) (roles []Role, err error)
	GetUsersOfRole(role Role) (users []User, err error)
	GetClubsOfUser(user User) (roles []Role, err error)
	RemoveUserFromClub(role Role) (users []User, err error)
}

func User_Club_App(domainUserRoleService domainservice.DomainUserRoleService) UserRoleApp {
	return &UserRoleImpl{
		domainUserRoleService: domainUserRoleService,
	}
}

func (u UserRoleImpl) AssignUserToRole(role Role) (users []User, err error) {
	domainRole := MapRoleAppToDomain(role)
	domainUsers, err := u.domainUserRoleService.AssignUserToRole(domainRole)
	if err != nil {
		return users, err
	}
	return MapUsersApp(domainUsers), err
}

func (u UserRoleImpl) AssignClubsToUser(user User) (roles []Role, err error) {
	panic("implement me")
}

func (u UserRoleImpl) GetUsersOfRole(role Role) (users []User, err error) {
	domainRole := MapRoleAppToDomain(role)
	 domainUsers, err := u.domainUserRoleService.GetUsersOfRole(domainRole)
	 if err!=nil {
		 return nil, err
	 }
	 return MapUsersApp(domainUsers), nil
}

func (u UserRoleImpl) GetClubsOfUser(user User) (roles []Role, err error) {
	panic("implement me")
}

func (u UserRoleImpl) RemoveUserFromClub(role Role) (users []User, err error) {
	panic("implement me")
}