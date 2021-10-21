package repo

import (
	"github.com/go-pg/pg/v10"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/postgress"
	"golang/homeworks/lesson10/util"
)

type UserRepoPostgressCreate interface {
	CreateUser(user entities.User) (userCreated *entities.User, err error)
}

func User_RepoPostgress_Create(postgressDB *postgress.PostgressDB) UserRepoPostgressCreate {
	return &UserRepoPostgressImpl{
		PostgressDB: postgressDB,
	}
}


func (u *UserRepoPostgressImpl) CreateUser(user entities.User) (userCreated *entities.User, err error) {
	var transaction *pg.Tx
	transaction, err  = u.PostgressDB.DB.Begin()
	if err != nil {
		return nil, err
	}
	_, err = transaction.Model(&user).Insert()
	if !util.Check_err(err, transaction) {
		return nil, err
	}

	for _, role := range user.Int_roles {
		user_role := entities.User_Role{
			User_id: user.Id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !util.Check_err(err, transaction) {
			return nil, err
		}
	}
	if err = transaction.Commit(); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}