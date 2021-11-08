package repo

import (
	"errors"
	"golang/homeworks/lesson10/entities"
"golang/homeworks/lesson10/infrastucture/databases"
)

type MemberRepo interface {
	GetMembers() (members []entities.Member, err error)
	GetMember(memberParam entities.Member) (member entities.Member, err error)
	Create(member entities.Member) (entities.Member, error)
	Update(member entities.Member) (entities.Member, error)
	Delete(member entities.Member) (members []entities.Member, err error)
}


type MemberPostgressRepoImpl struct {
	PostgressDB *databases.PostgressDB
}


type MemberMySQLRepoImpl struct {
	gormDB *databases.MySQLDB
}

func MemberMySQLRepo(mysqlDB *databases.MySQLDB) MemberRepo {
	return &MemberMySQLRepoImpl{
		gormDB: mysqlDB,
	}
}

func (m MemberMySQLRepoImpl) GetMembers() (members []entities.Member, err error) {
	err = m.gormDB.DB.Find(&members).Error
	return members, err
}

func (m MemberMySQLRepoImpl) GetMember(memberParam entities.Member) (member entities.Member, err error) {
	err = m.gormDB.DB.First(&member, "id=?", memberParam.Id).Error
	return member, err
}

func (m MemberMySQLRepoImpl) Create(member entities.Member) (entities.Member, error) {
	transaction := m.gormDB.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	err := transaction.Create(&member).Error
	if err != nil {
		transaction.Rollback()
		return member, err
	}
	if err = transaction.Commit().Error; err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func (m MemberMySQLRepoImpl) Update(member entities.Member) (entities.Member, error) {
	_, err := m.GetMember(member)
	if err !=  nil {
		return member, errors.New("Có lỗi khi get user")
	}
	err = m.gormDB.DB.Save(&member).Error
	return member, err
}

func (m MemberMySQLRepoImpl) Delete(member entities.Member) (members []entities.Member, err error) {
	err = m.gormDB.DB.Where("id=?", member.Id).Or("name=?", member.Name).Delete(&member).Error
	if err != nil {
		return members, err
	}
	return m.GetMembers()
}