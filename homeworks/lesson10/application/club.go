package application

import (
	domainmodel "golang/homeworks/lesson10/domain/model"
	domainservice "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/util"
)

type Club struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
}

type ClubImpl struct {
	domainClubService domainservice.DomainClubService
}

type ClubApp interface {
	ClubCreate(club Club) (err error)
	GetClub(clubParam Club) (club Club, err error)
	GetClubs() (clubs []Club, err error)
	Update(clubParam Club)(Club, error)
}

func Club_App(domainClubService domainservice.DomainClubService) ClubApp {
	return &ClubImpl{
		domainClubService: domainClubService,
	}
}

func (c ClubImpl) ClubCreate(club Club) (err error) {
	domainClub := domainmodel.Club{
		Id: util.NewID(),
		Name: club.Name,
	}
	return c.domainClubService.Create(domainClub)
}

func (c ClubImpl) GetClub(clubParam Club) (club Club, err error) {
	panic("implement me")
}

func (c ClubImpl) GetClubs() (clubs []Club, err error) {
	domainClubs, err := c.domainClubService.GetUsers()
	if err != nil {
		return clubs, err
	}
	return MapClubsDomainToApp(domainClubs), err
}

func (c ClubImpl) Update(clubParam Club) (Club, error) {
	panic("implement me")
}

func MapClubDomainToApp(domainClub domainmodel.Club) Club  {
	return Club{
		Id: domainClub.Id,
		Name: domainClub.Name,
	}
}

func MapClubsDomainToApp(domainClubs []domainmodel.Club) (clubs []Club)  {
	for _, domainClub := range domainClubs {
		clubs = append(clubs, MapClubDomainToApp(domainClub))
	}
	return clubs
}