package application

import (
	domainservice "golang/homeworks/lesson12/domain/services"
)

type MemberClubImpl struct {
	domainMemberClubService domainservice.DomainMemberClubService
}

type MemberClubApp interface {
	AssignMembersToClub(club Club, paramMembers []Member) (members []Member, err error)
	AssignClubsToMember(member Member) (clubs []Club, err error)
	GetMembersOfClub(club Club) (members []Member, err error)
	GetClubsOfMember(member Member) (clubs []Club, err error)
	RemoveMemberFromClub(club Club) (members []Member, err error)
}

func Member_Club_App(domainMemberClubService domainservice.DomainMemberClubService) MemberClubApp {
	return &MemberClubImpl{
		domainMemberClubService: domainMemberClubService,
	}
}

func (m MemberClubImpl) AssignMembersToClub(club Club, paramMembers []Member) (members []Member, err error) {
	domainClub := MapClubAppToDomain(club)
	domainMembers := MapMembersAppToDomain(paramMembers)
	domainMembers, err = m.domainMemberClubService.AssignMembersToClub(domainClub, domainMembers)
	if err != nil {
		return members, err
	}
	return MapMembersApp(domainMembers), err
}

func (m MemberClubImpl) AssignClubsToMember(member Member) (clubs []Club, err error) {
	panic("implement me")
}

func (m MemberClubImpl) GetMembersOfClub(club Club) (members []Member, err error) {
	domainClub := MapClubAppToDomain(club)
	domainMembers, err := m.domainMemberClubService.GetMemberOfClub(domainClub)
	if err!=nil {
		 return nil, err
	}
	return MapMembersApp(domainMembers), nil
}

func (m MemberClubImpl) GetClubsOfMember(member Member) (clubs []Club, err error) {
	panic("implement me")
}

func (m MemberClubImpl) RemoveMemberFromClub(club Club) (members []Member, err error) {
	panic("implement me")
}