package application

import (
	domainmodel "golang/homeworks/lesson12/domain/model"
	domainservice "golang/homeworks/lesson12/domain/services"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/util"
)

type User struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string	`json:"email"`
	Mobile     string	`json:"mobile"`
	Clubs []entities.Club `json:"clubs"`
}

type UserImpl struct {
	domainUserService domainservice.DomainUserService
}


type UserApp interface {
	UserCreate(user User) (appUser User, err error)
	GetUser(user User) (users User, err error)
	GetUsers() (users []User, err error)
	Update(user User)(User, error)
	DeleteUser(user User) (users []User, err error)
}

func User_App(domainUserService domainservice.DomainUserService) UserApp {
	return &UserImpl{
		domainUserService: domainUserService,
	}
}

func (u UserImpl) UserCreate(user User) (appUser User,err error) {
	domainUser := domainmodel.User{
		Id: util.NewID(),
		Name: user.Name,
		Email: user.Email,
		Mobile: user.Mobile,
	}
	domainUser, err = u.domainUserService.Create(domainUser)
	appUser = MapUserApp(domainUser)
	return appUser, err
}

func (u UserImpl) GetUsers() (users []User, err error)  {
	domainUsers, err := u.domainUserService.GetUsers()
	if err != nil {
		return nil, err
	}
	return MapUsersApp(domainUsers), nil
}

func (u UserImpl) GetUser(user User) (users User, err error) {
	entitiesUser := MapUserAppToUserEntities(user)
	domainUser, err := u.domainUserService.GetUser(entitiesUser)
	if err != nil {
		return users, err
	}
	return MapUserApp(domainUser), nil
}

func (u UserImpl) Update(user User)(User, error)  {
	domainUser := MapUserAppToUserDomain(user)
	domainUser, err := u.domainUserService.Update(domainUser)
	user = MapUserApp(domainUser)
	return user, err
}

func (u UserImpl) DeleteUser(user User) (users []User, err error) {
	entitiesUser := MapUserAppToUserEntities(user)
	domainUsers, err := u.domainUserService.DeleteUser(entitiesUser)
	if err != nil {
		return nil, err
	}
	return MapUsersApp(domainUsers), nil
}


func MapUserApp(domainUser domainmodel.User) User  {
	user := User{
		Id: domainUser.Id,
		Name: domainUser.Name,
		Email: domainUser.Email,
		Mobile: domainUser.Mobile,
	}
	return user
}

func MapUserAppToUserEntities(userApp User) entities.User  {
	user := entities.User{
		Id: userApp.Id,
		Name: userApp.Name,
		Email: userApp.Email,
		Mobile: userApp.Mobile,
	}
	return user
}

func MapUserAppToUserDomain(userApp User) domainmodel.User  {
	user := domainmodel.User{
		Id: userApp.Id,
		Name: userApp.Name,
		Email: userApp.Email,
		Mobile: userApp.Mobile,
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
		}
		users = append(users, user)
	}
	return users
}