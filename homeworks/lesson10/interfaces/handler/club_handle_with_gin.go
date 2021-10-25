package handler

import (
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	"net/http"
)

func (club *Club)GClubCreate(c *gin.Context) {
	clubApp := application.Club{}
	err := c.ShouldBindJSON(&clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	err = club.ClubApp.ClubCreate(clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusNotFound, err.Error())
	}
	util.GResponse(c, http.StatusCreated, clubApp)
}


func (club *Club)GGetClubs(c *gin.Context)  {
	clubApp, err := club.ClubApp.GetClubs()
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, clubApp)
}

func (club *Club)GGetClub(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {

	}
	clubApp := application.Club{
		Id: id,
	}
	clubApp, err := club.ClubApp.GetClub(clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, clubApp)
}


func (club *Club)GUpdateClub(c *gin.Context) {
	clubApp := application.Club{}
	err := c.ShouldBindJSON(&clubApp)
	clubApp, err = club.ClubApp.Update(clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, clubApp)
}
