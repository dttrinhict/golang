package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/postgress"
)

type UserRepoPostgressRead interface {
	GetUser() (users []entities.User, err error)
	GetUserById(userID string) (user *entities.User, err error)
}

func User_RepoPostgress_Read(postgressDB *postgress.PostgressDB) UserRepoPostgressRead {
	return &UserRepoPostgressImpl{
		PostgressDB: postgressDB,
	}
}


func (u *UserRepoPostgressImpl) GetUser() (users []entities.User, err error) {
	err = u.PostgressDB.DB.Model(&users).Select()
	return users, err
}

func (u *UserRepoPostgressImpl) GetUserById(userID string) (user *entities.User, err error) {
	panic("implement me")
}
