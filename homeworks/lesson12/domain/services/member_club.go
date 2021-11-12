package services

import (
	domainmodel "golang/homeworks/lesson12/domain/model"
	"golang/homeworks/lesson12/infrastucture/repo"
)

type DomainMemberClubImpl struct {
	memberClub repo.MemberClub
}

type DomainMemberClubService interface {
	AssignMembersToClub(club domainmodel.Club, paramMembers []domainmodel.Member) (members []domainmodel.Member, err error)
	AssignClubsToMember(member domainmodel.Member) (clubs []domainmodel.Club, err error)
	GetMemberOfClub(club domainmodel.Club) (members []domainmodel.Member, err error)
	GetClubsOfMember(member domainmodel.Member) (clubs []domainmodel.Club, err error)
	RemoveMemberFromClub(club domainmodel.Club, membersParam []domainmodel.Member) (members []domainmodel.Member, err error)
	Count(object string) (interface{}, error)
}

func DomainMemberClub(memberClub repo.MemberClub) DomainMemberClubService {
	return DomainMemberClubImpl{
		memberClub: memberClub,
	}
}

func (d DomainMemberClubImpl) AssignMembersToClub(club domainmodel.Club, paramMembers []domainmodel.Member) (members []domainmodel.Member, err error) {
	entitiesClub := MapClubDomainToEntities(club)
	entitiesMembers := MapMembersDomainToEntities(paramMembers)
	entitiesMembers, err = d.memberClub.AssignMembersToClub(entitiesClub, entitiesMembers)
	if err != nil {return members, err}
	return MapMembersEntitiesToDomain(entitiesMembers), err
}

func (d DomainMemberClubImpl) AssignClubsToMember(member domainmodel.Member) (clubs []domainmodel.Club, err error) {
	entitiesMember := MapMemberDomainToEntities(member)
	entitiesClubs, err := d.memberClub.AssignClubsToMember(entitiesMember)
	if err != nil { return clubs, err}
	return MapClubsEntitiesToDomain(entitiesClubs), err
}

func (d DomainMemberClubImpl) GetMemberOfClub(club domainmodel.Club) (members []domainmodel.Member, err error) {
	entitiesClub := MapClubDomainToEntities(club)
	entitiesMembers, err := d.memberClub.GetMembersOfClub(entitiesClub)
	return MapMembersEntitiesToDomain(entitiesMembers), err
}

func (d DomainMemberClubImpl) GetClubsOfMember(member domainmodel.Member) (clubs []domainmodel.Club, err error) {
	entitiesMember := MapMemberDomainToEntities(member)
	entitiesClubs, err := d.memberClub.GetClubsOfMember(entitiesMember)
	if err != nil {
		return clubs, err
	}
	return MapClubsEntitiesToDomain(entitiesClubs), err
}

func (d DomainMemberClubImpl) RemoveMemberFromClub(club domainmodel.Club, membersParam []domainmodel.Member) (members []domainmodel.Member, err error) {
	entitiesClub := MapClubDomainToEntities(club)
	entitiesMembers := MapMembersDomainToEntities(membersParam)
	entitiesMembers, err = d.memberClub.RemoveMembersFromClub(entitiesClub, entitiesMembers)
	if err != nil {
		return members, err
	}
	return MapMembersEntitiesToDomain(entitiesMembers), err
}

func (d DomainMemberClubImpl) Count(object string) (interface{}, error) {
	return d.memberClub.Count(object)
}