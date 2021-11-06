package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

func User_Role_MySQL_Repo(mysqlDB *databases.MySQLDB) UserRole {
	return &UserRoleMySQLImpl{
		mySQLDB: mysqlDB,
	}
}

func (u UserRoleMySQLImpl) AssignUserToRole(role entities.Role) (users []entities.User, err error) {
	transaction := u.mySQLDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	for _, user := range role.Users {
		user_role := entities.User_Role{
			Role_id: role.Id,
			User_id: user.Id,
		}
		err := transaction.Create(&user_role).Error
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

func (u UserRoleMySQLImpl) AssignRolesToUser(user entities.User) (roles []entities.Role, err error) {
	panic("implement me")
}

func (u UserRoleMySQLImpl) GetUsersOfRole(role entities.Role) (users []*entities.User, err error) {
	err = u.mySQLDB.DB.Where("id=?", role.Id).Preload("Users").Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role.Users, err
}

func (u UserRoleMySQLImpl) GetRolesOfUser(user entities.User) (roles []*entities.Role, err error) {
	err = u.mySQLDB.DB.Where("id=?", user.Id).Preload("Roles").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user.Roles, err
}

func (u UserRoleMySQLImpl) RemoveUserFromRole(role entities.Role) (users []entities.User, err error) {
	panic("implement me")
}