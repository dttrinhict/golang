package services

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/repo"
)

type DomainClubImpl struct {
	ClubRepo repo.ClubRepo
}

type DomainClubService interface {
	Create(club domainmodel.Club) (err error)
	GetUser(clubParam entities.Club)(club domainmodel.Club, err error)
	GetUsers()(clubs []domainmodel.Club, err error)
	Update(clubParam domainmodel.Club)(domainmodel.Club, error)
}

func DomainClub(clubRepo repo.ClubRepo) DomainClubService {
	return DomainClubImpl{
		ClubRepo: clubRepo,
	}
}

func (d DomainClubImpl) Create(club domainmodel.Club) (err error) {
	entitiesClub := entities.Club{
		Id:         club.Id,
		Name:       club.Name,
	}
	_,err = d.ClubRepo.Create(entitiesClub)
	return err
}

func (d DomainClubImpl) GetUser(clubParam entities.Club) (club domainmodel.Club, err error) {
	panic("implement me")
}

func (d DomainClubImpl) GetUsers() (clubs []domainmodel.Club, err error) {
	entitiesClubs, err := d.ClubRepo.GetClubs()
	return MapClubsEntitiesToDomain(entitiesClubs), err

}

func (d DomainClubImpl) Update(clubParam domainmodel.Club) (domainmodel.Club, error) {
	panic("implement me")
}


func MapClubEntitiesToDomain(entitiesClub entities.Club) domainmodel.Club {
	return domainmodel.Club{
		Id: entitiesClub.Id,
		Name: entitiesClub.Name,
	}
}

func MapClubDomainToEntities(domainClub domainmodel.Club) entities.Club {
	return entities.Club{
		Id: domainClub.Id,
		Name: domainClub.Name,
		Users: domainClub.Users,
	}
}

func MapClubsEntitiesToDomain(entitiesClubs []entities.Club) (domainClubs []domainmodel.Club) {
	for _, entitiesClub := range entitiesClubs {
		domainClubs = append(domainClubs, MapClubEntitiesToDomain(entitiesClub))
	}
	return domainClubs
}