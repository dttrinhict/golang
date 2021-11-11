package repo

import (
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
)

type UserRole interface {
	AssignUserToRole(role entities.Role) (users []entities.User, err error)
	AssignRolesToUser(user entities.User) (clubs []entities.Role, err error)
	GetUsersOfRole(role entities.Role) (users []entities.User, err error)
	GetRolesOfUser(user entities.User) (clubs []entities.Role, err error)
	RemoveUserFromRole(role entities.Role) (users []entities.User, err error)
}

type UserRoleMySQLImpl struct {
	mySQLDB *databases.MySQLDB
}

type UserRolePostgressImpl struct {
	postgressDB *databases.PostgressDB
}

