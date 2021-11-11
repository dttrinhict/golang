package repo

import (
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
)

type RoleRepo interface {
	GetRoles() (roles []entities.Role, err error)
	GetRole(roleParam entities.Role) (role entities.Role, err error)
	Create(role entities.Role) (entities.Role, error)
	Update(role entities.Role) (entities.Role, error)
	Delete(role entities.Role) (entities.Role, error)
}

type RolePostgressRepoImpl struct {
	PostgressDB *databases.PostgressDB
}


type RoleMySQLRepoImpl struct {
	gormDB *databases.MySQLDB
}
