package repo

import (
	"golang/homeworks/lesson10/infrastucture/databases"
)

type UserPostgressRepoImpl struct {
	PostgressDB *databases.PostgressDB
}


type UserMySQLRepoImpl struct {
	gormDB *databases.MySQLDB
}