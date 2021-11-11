package repo

import (
	"errors"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
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
	err = c.gormDB.DB.Find(&club, "id=?", clubParam.Id).Error
	return club, err
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
	checkExsisted, err := c.CheckExsisted(club)
	if err != nil {
		return club, err
	}
	if !checkExsisted {
		return club, errors.New("the club is not exsisted")
	}
	err = c.gormDB.DB.Save(&club).Error
	return club, err
}

func (c ClubMySQLRepoImpl) Delete(club entities.Club) (entities.Club, error) {
	panic("implement me")
}


func (c ClubMySQLRepoImpl)CheckExsisted(club entities.Club) (bool, error) {
	clubResult, err := c.GetClub(club)
	if err !=  nil || clubResult.Id == "" {
		return false, errors.New("the club is not exsisted")
	}
	return true, nil
}