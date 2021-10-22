package repo

import (
	"golang/homeworks/lesson10/infrastucture/postgress"
)

type UserPostgressRepoImpl struct {
	PostgressDB *postgress.PostgressDB
}


type UserMySQLRepoImpl struct {
	PostgressDB *postgress.PostgressDB
}