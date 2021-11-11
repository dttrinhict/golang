package repo

import (
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
)

type MemberClub interface {
	AssignMembersToClub(club entities.Club, paramMembers []entities.Member) (members []entities.Member, err error)
	AssignClubsToMember(member entities.Member) (clubs []entities.Club, err error)
	GetMembersOfClub(club entities.Club) (members []entities.Member, err error)
	GetClubsOfMember(member entities.Member) (clubs []entities.Club, err error)
	RemoveMembersFromClub(club entities.Club, membersParam []entities.Member) (members []entities.Member, err error)
}

type MemberClubMySQLImpl struct {
	mySQLDB *databases.MySQLDB
}

type MemberClubPostgressImpl struct {
	postgressDB *databases.PostgressDB
}


func MemeberClubMySQLRepo(mysqlDB *databases.MySQLDB) MemberClub {
	return &MemberClubMySQLImpl{
		mySQLDB: mysqlDB,
	}
}

func (m MemberClubMySQLImpl) AssignMembersToClub(club entities.Club, paramMembers []entities.Member) (members []entities.Member, err error) {
	transaction := m.mySQLDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	for _, member := range paramMembers {
		member_club := entities.Member_Club{
			Club_id: club.Id,
			Member_id: member.Id,
		}
		err := transaction.Create(&member_club).Error
		if err != nil {
			transaction.Rollback()
			return members, err
		}
		members = append(members, member)
	}
	if err = transaction.Commit().Error; err != nil {
		return members, err
	} else {
		return members, nil
	}
}

func (m MemberClubMySQLImpl) AssignClubsToMember(member entities.Member) (clubs []entities.Club, err error) {
	transaction := m.mySQLDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	for _, club := range member.Clubs {
		member_club := entities.Member_Club{
			Club_id: club.Id,
			Member_id: member.Id,
		}
		err := transaction.Create(&member_club).Error
		if err != nil {
			transaction.Rollback()
			return clubs, err
		}
		clubs = append(clubs, club)
	}
	if err = transaction.Commit().Error; err != nil {
		return clubs, err
	} else {
		return clubs, nil
	}
}

func (m MemberClubMySQLImpl) GetMembersOfClub(club entities.Club) (members []entities.Member, err error) {
	err = m.mySQLDB.DB.Where("id=?", club.Id).Preload("Members").Find(&club).Error
	return club.Members, err
}

func (m MemberClubMySQLImpl) GetClubsOfMember(member entities.Member) (clubs []entities.Club, err error) {
	err = m.mySQLDB.DB.Where("id=?", member.Id).Preload("Clubs").Find(&member).Error
	return member.Clubs, err
}

func (m MemberClubMySQLImpl) RemoveMembersFromClub(club entities.Club, membersParam []entities.Member) (members []entities.Member, err error) {
	memberClubs := []entities.Member_Club{}
	for _, member := range membersParam {
		memberClub := entities.Member_Club{
			Member_id: member.Id,
			Club_id: club.Id,
		}
		memberClubs = append(memberClubs, memberClub)
	}
	err =  m.mySQLDB.DB.Delete(memberClubs).Error
	if err != nil {
		return nil, err
	}
	return m.GetMembersOfClub(club)
}