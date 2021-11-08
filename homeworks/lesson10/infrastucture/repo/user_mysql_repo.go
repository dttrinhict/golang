package repo

import (
	"errors"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

func User_MySQL_Repo(mysqlDB *databases.MySQLDB) UserRepo {
	return &UserMySQLRepoImpl{
		gormDB: mysqlDB,
	}
}

func (u *UserMySQLRepoImpl) GetUers() (users []entities.User, err error) {
	err = u.gormDB.DB.Find(&users).Error
	return users, err
}

func (u *UserMySQLRepoImpl) GetUser(userParam entities.User) (user entities.User, err error) {
	user.Id = userParam.Id
	err = u.gormDB.DB.First(&user, "id=?", userParam.Id).Error
	return user, err
}

func (u *UserMySQLRepoImpl) Create(user entities.User) (entities.User, error) {
	transaction := u.gormDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	err := transaction.Create(&user).Error
	if err != nil {
		transaction.Rollback()
		return user, err
	}
	if err = transaction.Commit().Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (u *UserMySQLRepoImpl) Update(user entities.User) (entities.User, error) {
	_, err := u.GetUser(user)
	if err !=  nil {
		return user, errors.New("Có lỗi khi get user")
	}
	err = u.gormDB.DB.Save(&user).Error
	return user, err
}

func (u *UserMySQLRepoImpl) Delete(user entities.User) (users []entities.User, err error) {
	err = u.gormDB.DB.Where("id=?",user.Id).Or("name=?",user.Name).Delete(&user).Error
	if err != nil {
		return users, err
	}
	return u.GetUers()
}