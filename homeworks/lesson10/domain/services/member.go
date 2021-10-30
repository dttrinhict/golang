package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainMemberImpl struct {
	memberRepo repo.MemberRepo
}

type DomainMemberService interface {
	Create(member domainmodel.Member) (err error)
	GetMember(memberParam entities.Member)(member domainmodel.Member, err error)
	GetMembers()(members []domainmodel.Member, err error)
	Update(memberParam domainmodel.Member)(domainmodel.Member, error)
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
	return MemberMap(entitiesUser), err
}

func (m DomainMemberImpl) GetMembers() (members []domainmodel.Member, err error) {
	entitiesMembers, err := m.memberRepo.GetMembers()
	if err != nil {
		return nil, err
	}
	for _, entitiesMember := range entitiesMembers {
		members = append(members, MemberMap(entitiesMember))
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

func MemberMap(entitiesMember entities.Member) domainmodel.Member {
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

func MapMemberEntitiesToDomain(entitiesMembers []entities.Member) (domainMembers []domainmodel.Member) {
	for _, member := range entitiesMembers {
		domainMember := domainmodel.Member{
			Id: member.Id,
			Name: member.Name,
			Email: member.Email,
			Mobile: member.Mobile,
		}
		domainMembers = append(domainMembers, domainMember)
	}
	return domainMembers
}

func MapMembersEntitiesToDomain1(entitiesMembers []*entities.Member) (domainMembers []domainmodel.Member) {
	for _, member := range entitiesMembers {
		domainMember := domainmodel.Member{
			Id: member.Id,
			Name: member.Name,
			Email: member.Email,
			Mobile: member.Mobile,
		}
		domainMembers = append(domainMembers, domainMember)
	}
	return domainMembers
}