package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainUserRoleImpl struct {
	userRole repo.UserRole
}

type DomainUserRoleService interface {
	AssignUserToRole(role domainmodel.Role) (users []domainmodel.User, err error)
	AssignRolesToUser(user domainmodel.User) (clubs []domainmodel.Role, err error)
	GetUsersOfRole(role domainmodel.Role) (users []domainmodel.User, err error)
	GetRolesOfUser(user domainmodel.User) (clubs []domainmodel.Role, err error)
	RemoveUserFromRole(role domainmodel.Role) (users []domainmodel.User, err error)
}

func DomainUserClub(userRole repo.UserRole) DomainUserRoleService {
	return DomainUserRoleImpl{
		userRole: userRole,
	}
}

func (d DomainUserRoleImpl) AssignUserToRole(role domainmodel.Role) (users []domainmodel.User, err error) {
	entitiesClub := MapRoleDomainToEntities(role)
	entitiesUsers, err := d.userRole.AssignUserToRole(entitiesClub)
	if err != nil {return users, err}
	return MapUsersEntitiesToDomain(entitiesUsers), err
}

func (d DomainUserRoleImpl) AssignRolesToUser(user domainmodel.User) (roles []domainmodel.Role, err error) {
	panic("implement me")
}

func (d DomainUserRoleImpl) GetUsersOfRole(role domainmodel.Role) (users []domainmodel.User, err error) {
	entitiesRole := MapRoleDomainToEntities(role)
	entitiesUsers, err := d.userRole.GetUsersOfRole(entitiesRole)
	return MapUsersEntitiesToDomain1(entitiesUsers), err
}

func (d DomainUserRoleImpl) GetRolesOfUser(user domainmodel.User) (roles []domainmodel.Role, err error) {
	panic("implement me")
}

func (d DomainUserRoleImpl) RemoveUserFromRole(role domainmodel.Role) (users []domainmodel.User, err error) {
	panic("implement me")
}