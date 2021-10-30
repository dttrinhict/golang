package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

type UserClub interface {
	AssignUserToClub(club entities.Club) (users []entities.User, err error)
	AssignClubsToUser(user entities.User) (clubs []entities.Club, err error)
	GetUsersOfClub(club entities.Club) (users []*entities.User, err error)
	GetClubsOfUser(user entities.User) (clubs []*entities.Club, err error)
	RemoveUserFromClub(club entities.Club) (users []entities.User, err error)
}

type UserClubMySQLImpl struct {
	mySQLDB *databases.MySQLDB
}

type UserClubPostgressImpl struct {
	postgressDB *databases.PostgressDB
}

