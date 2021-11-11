package services

import (
	domainmodel "golang/homeworks/lesson12/domain/model"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/repo"
)

type DomainMemberImpl struct {
	memberRepo repo.MemberRepo
}

type DomainMemberService interface {
	Create(member domainmodel.Member) (err error)
	GetMember(memberParam entities.Member)(member domainmodel.Member, err error)
	GetMembers()(members []domainmodel.Member, err error)
	Update(memberParam domainmodel.Member)(domainmodel.Member, error)
	Delete(memberParam domainmodel.Member)([]domainmodel.Member, error)
}

func DomainMember(memberRepo repo.MemberRepo) DomainMemberService {
	return DomainMemberImpl{
		memberRepo: memberRepo,
	}
}


func (m DomainMemberImpl) Create(member domainmodel.Member) (err error) {
	memberStorePostgress := entities.Member{
		Id:         member.Id,
		Name:       member.Name,
		Email:      member.Email,
		Mobile:     member.Mobile,
	}
	_,err = m.memberRepo.Create(memberStorePostgress)
	return err
}

func (m DomainMemberImpl) GetMember(memberParam entities.Member) (member domainmodel.Member, err error) {
	entitiesUser, err := m.memberRepo.GetMember(memberParam)
	if err != nil {
		return domainmodel.Member{}, err
	}
	return MapMemberEntitiesToDomain(entitiesUser), err
}

func (m DomainMemberImpl) GetMembers() (members []domainmodel.Member, err error) {
	entitiesMembers, err := m.memberRepo.GetMembers()
	if err != nil {
		return nil, err
	}
	for _, entitiesMember := range entitiesMembers {
		members = append(members, MapMemberEntitiesToDomain(entitiesMember))
	}
	return members, err
}

func (m DomainMemberImpl) Update(memberParam domainmodel.Member)(domainmodel.Member, error) {
	entitiesMember := MapMemberDomainToEntities(memberParam)
	_, err := m.memberRepo.Update(entitiesMember)
	if err != nil {
		return memberParam, err
	}
	return memberParam, nil
}

func (m DomainMemberImpl) Delete(memberParam domainmodel.Member)([]domainmodel.Member, error) {
	entitiesMember := MapMemberDomainToEntities(memberParam)
	_, err := m.memberRepo.Delete(entitiesMember)
	if err != nil {
		return nil, err
	}
	return m.GetMembers()
}

func MapMemberEntitiesToDomain(entitiesMember entities.Member) domainmodel.Member {
	domainUser := domainmodel.Member{
		Id: entitiesMember.Id,
		Name: entitiesMember.Name,
		Email: entitiesMember.Email,
		Mobile: entitiesMember.Mobile,
	}
	return domainUser
}

func MapMemberDomainToEntities(domainMember domainmodel.Member) entities.Member {
	entitiesMember := entities.Member{
		Id: domainMember.Id,
		Name: domainMember.Name,
		Email: domainMember.Email,
		Mobile: domainMember.Mobile,
	}
	return entitiesMember
}

func MapMembersEntitiesToDomain(entitiesMembers []entities.Member) (domainMembers []domainmodel.Member) {
	for _, entitiesMember := range entitiesMembers {
		domainMembers = append(domainMembers, MapMemberEntitiesToDomain(entitiesMember))
	}
	return domainMembers
}

func MapMembersDomainToEntities(domainMembers []domainmodel.Member) (entitiesMembers []entities.Member) {
		for _, domainMember := range domainMembers {
			entitiesMembers = append(entitiesMembers,MapMemberDomainToEntities(domainMember))
		}
	return entitiesMembers
}