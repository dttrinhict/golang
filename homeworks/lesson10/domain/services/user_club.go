package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainUserClubImpl struct {
	userClub repo.UserClub
}

type DomainUserClubService interface {
	AssignUserToClub(club domainmodel.Club) (users []domainmodel.User, err error)
	AssignClubsToUser(user domainmodel.User) (clubs []domainmodel.Club, err error)
	GetUsersOfClub(club domainmodel.Club) (users []domainmodel.User, err error)
	GetClubsOfUser(user domainmodel.User) (clubs []domainmodel.Club, err error)
	RemoveUserFromClub(club domainmodel.Club) (users []domainmodel.User, err error)
}

func DomainUserClub(userClub repo.UserClub) DomainUserClubService {
	return DomainUserClubImpl{
		userClub: userClub,
	}
}

func (d DomainUserClubImpl) AssignUserToClub(club domainmodel.Club) (users []domainmodel.User, err error) {
	entitiesClub := MapClubDomainToEntities(club)
	entitiesUsers, err := d.userClub.AssignUserToClub(entitiesClub)
	if err != nil {return users, err}
	return MapUsersEntitiesToDomain(entitiesUsers), err
}

func (d DomainUserClubImpl) AssignClubsToUser(user domainmodel.User) (clubs []domainmodel.Club, err error) {
	panic("implement me")
}

func (d DomainUserClubImpl) GetUsersOfClub(club domainmodel.Club) (users []domainmodel.User, err error) {
	entitiesClub := MapClubDomainToEntities(club)
	entitiesUsers, err := d.userClub.GetUsersOfClub(entitiesClub)
	return MapUsersEntitiesToDomain1(entitiesUsers), err
}

func (d DomainUserClubImpl) GetClubsOfUser(user domainmodel.User) (clubs []domainmodel.Club, err error) {
	panic("implement me")
}

func (d DomainUserClubImpl) RemoveUserFromClub(club domainmodel.Club) (users []domainmodel.User, err error) {
	panic("implement me")
}