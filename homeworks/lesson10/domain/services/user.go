package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainUserImpl struct {
	userRepo repo.UserRepo
}

type DomainUserService interface {
	Create(user domainmodel.User) (err error)
	GetUser(userParam entities.User)(user domainmodel.User, err error)
	GetUsers()(users []domainmodel.User, err error)
}

func DomainUser(userRepo repo.UserRepo) DomainUserService {
	return DomainUserImpl{
		userRepo: userRepo,
	}
}


func (u DomainUserImpl) Create(user domainmodel.User) (err error) {
	userStorePostgress := entities.User{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Mobile:     user.Mobile,
		Int_roles:  user.Int_roles,
		Enum_roles: user.Enum_roles,
	}
	_,err = u.userRepo.Create(userStorePostgress)
	return err
}

func (u DomainUserImpl) GetUser(userParam entities.User) (user domainmodel.User, err error) {
	entitiesUser, err := u.userRepo.GetUser(userParam)
	if err != nil {
		return domainmodel.User{}, err
	}
	return UserMap(entitiesUser), err
}

func (u DomainUserImpl) GetUsers() (users []domainmodel.User, err error) {
	entitiesUsers, err := u.userRepo.GetUers()
	if err != nil {
		return nil, err
	}
	for _, entitiesUser := range entitiesUsers {
		users = append(users, UserMap(entitiesUser))
	}
	return users, err
}


func UserMap(entitiesUser entities.User) domainmodel.User {
	domainUser := domainmodel.User{
		Id: entitiesUser.Id,
		Name: entitiesUser.Name,
		Email: entitiesUser.Email,
		Mobile: entitiesUser.Mobile,
		Int_roles: entitiesUser.Int_roles,
		Enum_roles: entitiesUser.Enum_roles,
	}
	return domainUser
}