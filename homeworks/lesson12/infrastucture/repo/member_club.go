package repo

import (
	"errors"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
	"strings"
)

type MemberClub interface {
	AssignMembersToClub(club entities.Club, paramMembers []entities.Member) (members []entities.Member, err error)
	AssignClubsToMember(member entities.Member) (clubs []entities.Club, err error)
	GetMembersOfClub(club entities.Club) (members []entities.Member, err error)
	GetClubsOfMember(member entities.Member) (clubs []entities.Club, err error)
	RemoveMembersFromClub(club entities.Club, membersParam []entities.Member) (members []entities.Member, err error)
	Count(object string) (interface{}, error)
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

func (m MemberClubMySQLImpl) Count(object string) (interface{}, error) {
	if strings.Compare(object, "members-of-clubs") == 0 {
		return m.CountMembersOfClubs()
	}
	if strings.Compare(object, "clubs-of-members") == 0 {
		return m.CountClubsOfMembers()
	}
	return nil, errors.New("Wrong input data")
}

func (m MemberClubMySQLImpl) CountMembersOfClubs() (interface{}, error) {
	var queryResult []struct {
		Name          string
		NumberMembers int
	}
	err := m.mySQLDB.DB.Raw("SELECT `clubs`.`name` as Name, COUNT(`members`.`id`) as NumberMembers FROM ((`member_club` INNER JOIN `clubs` ON `member_club`.`club_id` = `clubs`.`id`) INNER JOIN `members` ON `member_club`.`member_id` = `members`.`id`) GROUP BY name").Scan(&queryResult).Error
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}

func (m MemberClubMySQLImpl) CountClubsOfMembers() (interface{}, error) {
	var queryResult []struct {
		Name          string
		NumberClubs int
	}
	err := m.mySQLDB.DB.Raw("SELECT `members`.`name` as Name, COUNT(`clubs`.`id`) as NumberClubs FROM ((`member_club` INNER JOIN `clubs` ON `member_club`.`club_id` = `clubs`.`id`) INNER JOIN `members` ON `member_club`.`member_id` = `members`.`id`) GROUP BY name").Scan(&queryResult).Error
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}
