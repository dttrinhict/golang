package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainUserImpl struct {
	userRepoPostgressCreate repo.UserRepoPostgressCreate
	userRepoPostgressRead repo.UserRepoPostgressRead
}



type DomainUserService interface {
	UserCreate(user *domainmodel.User) (err error)
	GetUser()(user []domainmodel.User, err error)
}

func DomainUser(userRepoPostgressCreate repo.UserRepoPostgressCreate, userRepoPostgressRead repo.UserRepoPostgressRead) DomainUserService {
	return DomainUserImpl{
		userRepoPostgressCreate: userRepoPostgressCreate,
		userRepoPostgressRead: userRepoPostgressRead,
	}
}


func (u DomainUserImpl) UserCreate(user *domainmodel.User) (err error) {
	userStorePostgress := entities.User{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Mobile:     user.Mobile,
		Int_roles:  user.Int_roles,
		Enum_roles: user.Enum_roles,
	}
	_,err = u.userRepoPostgressCreate.CreateUser(userStorePostgress)
	return err
}

func (u DomainUserImpl) GetUser() (users []domainmodel.User, err error) {
	entitiesUsers, err := u.userRepoPostgressRead.GetUser()
	if err != nil {
		return nil, err
	}
	for _, entitiesUser := range entitiesUsers {
		users = append(users, UserMap(entitiesUser))
	}
	return users, nil
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