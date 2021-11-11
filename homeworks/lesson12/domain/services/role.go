package services

import (
	domainmodel "golang/homeworks/lesson12/domain/model"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/repo"
)

type DomainRoleImpl struct {
	RoleRepo repo.RoleRepo
}

type DomainRoleService interface {
	Create(role domainmodel.Role) (err error)
	GetUser(roleParam entities.Role)(role domainmodel.Role, err error)
	GetUsers()(roles []domainmodel.Role, err error)
	Update(roleParam domainmodel.Role)(domainmodel.Role, error)
}

func DomainRole(roleRepo repo.RoleRepo) DomainRoleService {
	return DomainRoleImpl{
		RoleRepo: roleRepo,
	}
}

func (d DomainRoleImpl) Create(role domainmodel.Role) (err error) {
	entitiesRole := entities.Role{
		Id:         role.Id,
		Name:       role.Name,
	}
	_,err = d.RoleRepo.Create(entitiesRole)
	return err
}

func (d DomainRoleImpl) GetUser(roleParam entities.Role) (role domainmodel.Role, err error) {
	panic("implement me")
}

func (d DomainRoleImpl) GetUsers() (roles []domainmodel.Role, err error) {
	entitiesRoles, err := d.RoleRepo.GetRoles()
	return MapRolesEntitiesToDomain(entitiesRoles), err

}

func (d DomainRoleImpl) Update(roleParam domainmodel.Role) (domainmodel.Role, error) {
	panic("implement me")
}


func MapRoleEntitiesToDomain(entitiesRole entities.Role) domainmodel.Role {
	return domainmodel.Role{
		Id: entitiesRole.Id,
		Name: entitiesRole.Name,
	}
}

func MapRoleDomainToEntities(domainRole domainmodel.Role) entities.Role {
	return entities.Role{
		Id: domainRole.Id,
		Name: domainRole.Name,
		Users: domainRole.Users,
	}
}

func MapRolesEntitiesToDomain(entitiesRoles []entities.Role) (domainRoles []domainmodel.Role) {
	for _, entitiesRole := range entitiesRoles {
		domainRoles = append(domainRoles, MapRoleEntitiesToDomain(entitiesRole))
	}
	return domainRoles
}