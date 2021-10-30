package repo

import (
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
)

type MemberClub interface {
	AssignMemberToClub(club entities.Club) (members []entities.Member, err error)
	AssignClubsToMember(member entities.Member) (clubs []entities.Club, err error)
	GetMembersOfClub(club entities.Club) (members []*entities.Member, err error)
	GetClubsOfMember(member entities.Member) (clubs []*entities.Club, err error)
	RemoveMemberFromClub(club entities.Club) (members []entities.Member, err error)
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

func (m MemberClubMySQLImpl) AssignMemberToClub(club entities.Club) (members []entities.Member, err error) {
	transaction := m.mySQLDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	for _, member := range club.Members {
		member_club := entities.Member_Club{
			Club_id: club.Id,
			Member_id: member.Id,
		}
		err := transaction.Create(&member_club).Error
		if err != nil {
			transaction.Rollback()
			return members, err
		}
		members = append(members, *member)
	}
	if err = transaction.Commit().Error; err != nil {
		return members, err
	} else {
		return members, nil
	}
}

func (m MemberClubMySQLImpl) AssignClubsToMember(member entities.Member) (clubs []entities.Club, err error) {
	panic("implement me")
}

func (m MemberClubMySQLImpl) GetMembersOfClub(club entities.Club) (members []*entities.Member, err error) {
	panic("implement me")
}

func (m MemberClubMySQLImpl) GetClubsOfMember(member entities.Member) (clubs []*entities.Club, err error) {
	panic("implement me")
}

func (m MemberClubMySQLImpl) RemoveMemberFromClub(club entities.Club) (members []entities.Member, err error) {
	panic("implement me")
}