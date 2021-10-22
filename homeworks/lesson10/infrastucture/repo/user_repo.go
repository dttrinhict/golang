package repo

import "golang/homeworks/lesson10/entities"

type UserRepo interface {
	GetUers() (users []entities.User, err error)
	GetUser(userParam *entities.User) (user *entities.User, err error)
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(user *entities.User) (*entities.User, error)
}
