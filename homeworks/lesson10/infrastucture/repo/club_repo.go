package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

type ClubRepo interface {
	GetClubs() (clubs []entities.Club, err error)
	GetClub(clubParam entities.Club) (club entities.Club, err error)
	Create(club entities.Club) (entities.Club, error)
	Update(club entities.Club) (entities.Club, error)
	Delete(club entities.Club) (entities.Club, error)
}

type ClubPostgressRepoImpl struct {
	PostgressDB *databases.PostgressDB
}


type ClubMySQLRepoImpl struct {
	gormDB *databases.MySQLDB
}
