package handler

import (
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	"net/http"
)

func (userClub *UserClub)GAssignUserToClub(c *gin.Context) {
	clubApp := application.Club{}
	err := c.ShouldBindJSON(&clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	users, err := userClub.UserClubApp.AssignUserToClub(clubApp)
	if err != nil {
		util.GResponseErr(c, http.StatusNotFound, err.Error())
	}
	util.GResponse(c, http.StatusCreated, users)
}
