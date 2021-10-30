package repo

import (
"golang/homeworks/lesson10/entities"
"golang/homeworks/lesson10/infrastucture/databases"
)

type MemberRepo interface {
	GetMembers() (members []entities.Member, err error)
	GetMember(memberParam entities.Member) (member entities.Member, err error)
	Create(member entities.Member) (entities.Member, error)
	Update(member entities.Member) (entities.Member, error)
	Delete(member entities.Member) (entities.Member, error)
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
	panic("implement me")
}

func (m MemberMySQLRepoImpl) GetMember(memberParam entities.Member) (member entities.Member, err error) {
	panic("implement me")
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
	panic("implement me")
}

func (m MemberMySQLRepoImpl) Delete(member entities.Member) (entities.Member, error) {
	panic("implement me")
}