package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

func Club_MySQL_Repo(mysqlDB *databases.MySQLDB) ClubRepo {
	return &ClubMySQLRepoImpl{
		gormDB: mysqlDB,
	}
}

func (c ClubMySQLRepoImpl) GetClubs() (clubs []entities.Club, err error) {
	err = c.gormDB.DB.Find(&clubs).Error
	return clubs, err
}

func (c ClubMySQLRepoImpl) GetClub(clubParam entities.Club) (club entities.Club, err error) {
	panic("implement me")
}

func (c ClubMySQLRepoImpl) Create(club entities.Club) (entities.Club, error) {
	transaction := c.gormDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	err := transaction.Create(&club).Error
	if err != nil {
		transaction.Rollback()
		return club, err
	}
	if err = transaction.Commit().Error; err != nil {
		return club, err
	} else {
		return club, nil
	}
}

func (c ClubMySQLRepoImpl) Update(club entities.Club) (entities.Club, error) {
	panic("implement me")
}

func (c ClubMySQLRepoImpl) Delete(club entities.Club) (entities.Club, error) {
	panic("implement me")
}
