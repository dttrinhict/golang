package repo

import (
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
)

type UserRepo interface {
	GetUers() (users []entities.User, err error)
	GetUser(userParam entities.User) (user entities.User, err error)
	Create(user entities.User) (entities.User, error)
	Update(user entities.User) (entities.User, error)
	Delete(user entities.User) (users []entities.User, err error)
}


type UserPostgressRepoImpl struct {
	PostgressDB *databases.PostgressDB
}


type UserMySQLRepoImpl struct {
	gormDB *databases.MySQLDB
}
