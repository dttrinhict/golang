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
	Create(user domainmodel.User) (domainUser domainmodel.User, err error)
	GetUser(userParam entities.User)(user domainmodel.User, err error)
	GetUsers()(users []domainmodel.User, err error)
	Update(userParam domainmodel.User)(domainmodel.User, error)
	DeleteUser(userParam entities.User) (users []domainmodel.User, err error)
}

func DomainUser(userRepo repo.UserRepo) DomainUserService {
	return DomainUserImpl{
		userRepo: userRepo,
	}
}


func (u DomainUserImpl) Create(user domainmodel.User) (domainUser domainmodel.User, err error) {
	userStorePostgress := entities.User{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Mobile:     user.Mobile,
	}
	entitiesUser ,err := u.userRepo.Create(userStorePostgress)
	domainUser = MapUserEntitiesToDomain(entitiesUser)
	return domainUser, err
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

func (u DomainUserImpl) Update(userParam domainmodel.User)(domainmodel.User, error) {
	entitiesUser := MapUserDomainToEntities(userParam)
	entitiesUser, err := u.userRepo.Update(entitiesUser)
	if err != nil {
		return userParam, err
	}
	userParam = MapUserEntitiesToDomain(entitiesUser)
	return userParam, nil
}

func (u DomainUserImpl) DeleteUser(userParam entities.User) (users []domainmodel.User, err error) {
	entitiesUsers, err := u.userRepo.Delete(userParam)
	if err != nil {
		return nil, err
	}
	return MapUsersEntitiesToDomain(entitiesUsers), err
}

func UserMap(entitiesUser entities.User) domainmodel.User {
	domainUser := domainmodel.User{
		Id: entitiesUser.Id,
		Name: entitiesUser.Name,
		Email: entitiesUser.Email,
		Mobile: entitiesUser.Mobile,
	}
	return domainUser
}

func MapUserDomainToEntities(domainUser domainmodel.User) entities.User {
	entitiesUser := entities.User{
		Id: domainUser.Id,
		Name: domainUser.Name,
		Email: domainUser.Email,
		Mobile: domainUser.Mobile,
	}
	return entitiesUser
}

func MapUserEntitiesToDomain(entitiesUser entities.User) (domainUsers domainmodel.User) {
		domainUser := domainmodel.User{
			Id: entitiesUser.Id,
			Name: entitiesUser.Name,
			Email: entitiesUser.Email,
			Mobile: entitiesUser.Mobile,
		}
	return domainUser
}

func MapUsersEntitiesToDomain(entitiesUsers []entities.User) (domainUsers []domainmodel.User) {
	for _, user := range entitiesUsers {
		domainUser := domainmodel.User{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Mobile: user.Mobile,
		}
		domainUsers = append(domainUsers, domainUser)
	}
	return domainUsers
}

func MapUsersEntitiesToDomain1(entitiesUsers []*entities.User) (domainUsers []domainmodel.User) {
	for _, user := range entitiesUsers {
		domainUser := domainmodel.User{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Mobile: user.Mobile,
		}
		domainUsers = append(domainUsers, domainUser)
	}
	return domainUsers
}