package application

import (
	domainmodel "golang/homeworks/lesson12/domain/model"
	domainservice "golang/homeworks/lesson12/domain/services"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/util"
)

type Member struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string	`json:"email"`
	Mobile     string	`json:"mobile"`
	Clubs []entities.Club `json:"clubs"`
}

type MemberImpl struct {
	domainMemberService domainservice.DomainMemberService
}


type MemberApp interface {
	Create(member Member) (err error)
	GetMember(memberParam Member) (member Member, err error)
	GetMembers() (members []Member, err error)
	Update(member Member)(Member, error)
	Delete(member Member)([]Member, error)
}

func Member_App(domainMemberService domainservice.DomainMemberService) MemberApp {
	return &MemberImpl{
		domainMemberService: domainMemberService,
	}
}

func (m MemberImpl) Create(member Member) (err error) {
	domainMember := domainmodel.Member{
		Id: util.NewID(),
		Name: member.Name,
		Email: member.Email,
		Mobile: member.Mobile,
	}
	return m.domainMemberService.Create(domainMember)
}

func (m MemberImpl) GetMembers() (members []Member, err error)  {
	domainMembers, err := m.domainMemberService.GetMembers()
	if err != nil {
		return nil, err
	}
	return MapMembersApp(domainMembers), nil
}

func (m MemberImpl) GetMember(member Member) (members Member, err error) {
	entitiesMember := MapMemberAppToEntities(member)
	domainMember, err := m.domainMemberService.GetMember(entitiesMember)
	if err != nil {
		return members, err
	}
	return MapMemberApp(domainMember), nil
}

func (m MemberImpl) Update(member Member)(Member, error)  {
	domainMember := MapMemberAppToDomain(member)
	_, err := m.domainMemberService.Update(domainMember)
	return member, err
}

func (m MemberImpl) Delete(member Member)([]Member, error)  {
	domainMember := MapMemberAppToDomain(member)
	_, err := m.domainMemberService.Delete(domainMember)
	return nil, err
}

func MapMemberApp(domainMember domainmodel.Member) Member  {
	member := Member{
		Id: domainMember.Id,
		Name: domainMember.Name,
		Email: domainMember.Email,
		Mobile: domainMember.Mobile,
	}
	return member
}

func MapMemberAppToEntities(memberApp Member) entities.Member  {
	member := entities.Member{
		Id: memberApp.Id,
		Name: memberApp.Name,
		Email: memberApp.Email,
		Mobile: memberApp.Mobile,
	}
	return member
}

func MapMemberAppToDomain(memberApp Member) domainmodel.Member  {
	member := domainmodel.Member{
		Id: memberApp.Id,
		Name: memberApp.Name,
		Email: memberApp.Email,
		Mobile: memberApp.Mobile,
	}
	return member
}

func MapMembersAppToDomain(membersApp []Member) (domainMembers[]domainmodel.Member)  {
	for _, memberApp := range membersApp {
		domainMembers = append(domainMembers, MapMemberAppToDomain(memberApp))
	}
	return domainMembers
}

func MapMembersApp(domainMembers []domainmodel.Member) (members []Member) {
	for _, domainMember := range domainMembers {
		member := Member{
			Id: domainMember.Id,
			Name: domainMember.Name,
			Email: domainMember.Email,
			Mobile: domainMember.Mobile,
		}
		members = append(members, member)
	}
	return members
}