package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

func User_Club_MySQL_Repo(mysqlDB *databases.MySQLDB) UserClub {
	return &UserClubMySQLImpl{
		mySQLDB: mysqlDB,
	}
}

func (u UserClubMySQLImpl) AssignUserToClub(club entities.Club) (users []entities.User, err error) {
	transaction := u.mySQLDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	for _, user := range club.Users {
		user_club := entities.User_Club{
			Club_id: club.Id,
			User_id: user.Id,
		}
		err := transaction.Create(&user_club).Error
		if err != nil {
			transaction.Rollback()
			return users, err
		}
		users = append(users, *user)
	}
	if err = transaction.Commit().Error; err != nil {
		return users, err
	} else {
		return users, nil
	}
}

func (u UserClubMySQLImpl) AssignClubsToUser(user entities.User) (clubs []entities.Club, err error) {
	panic("implement me")
}

func (u UserClubMySQLImpl) GetUsersOfClub(club entities.Club) (users []*entities.User, err error) {
	err = u.mySQLDB.DB.Where("id=?", club.Id).Preload("Users").Find(&club).Error
	if err != nil {
		return nil, err
	}
	return club.Users, err
}

func (u UserClubMySQLImpl) GetClubsOfUser(user entities.User) (clubs []*entities.Club, err error) {
	err = u.mySQLDB.DB.Where("id=?", user.Id).Preload("Clubs").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user.Clubs, err
}

func (u UserClubMySQLImpl) RemoveUserFromClub(club entities.Club) (users []entities.User, err error) {
	panic("implement me")
}