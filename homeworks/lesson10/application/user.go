package application

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	domainservice "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/util"
)

type User struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string	`json:"email"`
	Mobile     string	`json:"mobile"`
	Int_roles  []int    `json:"int_roles"`
	Enum_roles []string `json:"enum_roles"`
}

type UserImpl struct {
	domainUserService domainservice.DomainUserService
}


type UserApp interface {
	UserCreate(user User) (err error)
	GetUser() (users User, err error)
	GetUsers() (users []User, err error)
}

func User_App(domainUserService domainservice.DomainUserService) UserApp {
	return &UserImpl{
		domainUserService: domainUserService,
	}
}

func (u UserImpl) UserCreate(user User) (err error) {
	domainUser := domainmodel.User{
		Id: util.NewID(),
		Name: user.Name,
		Email: user.Email,
		Mobile: user.Mobile,
		Int_roles: user.Int_roles,
		Enum_roles: user.Enum_roles,
	}
	return u.domainUserService.Create(domainUser)
}

func (u UserImpl) GetUsers() (users []User, err error)  {
	domainUsers, err := u.domainUserService.GetUsers()
	if err != nil {
		return nil, err
	}
	return MapUsersApp(domainUsers), nil
}

func (u UserImpl) GetUser() (users User, err error) {
	panic("implement me")
}

func MapUserApp(domainUser domainmodel.User) User  {
	user := User{
		Id: domainUser.Id,
		Name: domainUser.Name,
		Email: domainUser.Email,
		Mobile: domainUser.Mobile,
		Int_roles: domainUser.Int_roles,
		Enum_roles: domainUser.Enum_roles,
	}
	return user
}

func MapUsersApp(domainUsers []domainmodel.User) (users []User) {
	for _, domainUser := range domainUsers {
		user := User{
			Id: domainUser.Id,
			Name: domainUser.Name,
			Email: domainUser.Email,
			Mobile: domainUser.Mobile,
			Int_roles: domainUser.Int_roles,
			Enum_roles: domainUser.Enum_roles,
		}
		users = append(users, user)
	}
	return users
}