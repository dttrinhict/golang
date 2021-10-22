package repo

import (
	"github.com/go-pg/pg/v10"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/postgress"
	"golang/homeworks/lesson10/util"
)

func User_Postgress_Repo(postgressDB *postgress.PostgressDB) UserRepo {
	return &UserPostgressRepoImpl{
		PostgressDB: postgressDB,
	}
}

func (u *UserPostgressRepoImpl) GetUers() (users []entities.User, err error) {
	err = u.PostgressDB.DB.Model(&users).Select()
	return users, err
}

func (u *UserPostgressRepoImpl) GetUser(userParam entities.User) (user entities.User, err error) {
	user.Id = userParam.Id
	err = u.PostgressDB.DB.Model(&user).WherePK().Select()
	return user, err
}

func (u *UserPostgressRepoImpl) Create(user entities.User) (entities.User, error) {
	var transaction *pg.Tx
	transaction, err  := u.PostgressDB.DB.Begin()
	if err != nil {
		return user, err
	}
	_, err = transaction.Model(&user).Insert()
	if !util.Check_err(err, transaction) {
		return user, err
	}

	for _, role := range user.Int_roles {
		user_role := entities.User_Role{
			User_id: user.Id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !util.Check_err(err, transaction) {
			return user, err
		}
	}
	if err = transaction.Commit(); err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (u *UserPostgressRepoImpl) Update(user entities.User) (entities.User, error) {
	panic("implement me")
}

func (u *UserPostgressRepoImpl) Delete(user entities.User) (entities.User, error) {
	panic("implement me")
}