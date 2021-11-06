package application

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	domainservice "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/util"
)

type Role struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Users []*entities.User `json:"users"`
}

type RoleImpl struct {
	domainRoleService domainservice.DomainRoleService
}

type RoleApp interface {
	RoleCreate(role Role) (err error)
	GetRole(roleParam Role) (role Role, err error)
	GetRoles() (roles []Role, err error)
	Update(roleParam Role)(Role, error)
}

func Role_App(domainRoleService domainservice.DomainRoleService) RoleApp {
	return &RoleImpl{
		domainRoleService: domainRoleService,
	}
}

func (c RoleImpl) RoleCreate(role Role) (err error) {
	domainRole := domainmodel.Role{
		Id: util.NewID(),
		Name: role.Name,
	}
	return c.domainRoleService.Create(domainRole)
}

func (c RoleImpl) GetRole(roleParam Role) (role Role, err error) {
	panic("implement me")
}

func (c RoleImpl) GetRoles() (roles []Role, err error) {
	domainRoles, err := c.domainRoleService.GetUsers()
	if err != nil {
		return roles, err
	}
	return MapRolesDomainToApp(domainRoles), err
}

func (c RoleImpl) Update(roleParam Role) (Role, error) {
	panic("implement me")
}

func MapRoleDomainToApp(domainRole domainmodel.Role) Role  {
	return Role{
		Id: domainRole.Id,
		Name: domainRole.Name,
		Users: domainRole.Users,
	}
}

func MapRoleAppToDomain(role Role) domainmodel.Role  {
	return domainmodel.Role{
		Id: role.Id,
		Name: role.Name,
		Users: role.Users,
	}
}

func MapRolesDomainToApp(domainRoles []domainmodel.Role) (roles []Role)  {
	for _, domainRole := range domainRoles {
		roles = append(roles, MapRoleDomainToApp(domainRole))
	}
	return roles
}