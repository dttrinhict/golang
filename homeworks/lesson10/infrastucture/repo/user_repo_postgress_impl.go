package repo

import "golang/homeworks/lesson10/infrastucture/postgress"

type UserRepoPostgressImpl struct {
	PostgressDB *postgress.PostgressDB
}
