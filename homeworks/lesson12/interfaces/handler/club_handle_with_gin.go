package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson12/application"
	"golang/homeworks/lesson12/util"
	ilogger "golang/homeworks/lesson12/util/logger"
)

func (club *Club)GClubCreate(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	clubApp := application.Club{}
	err := c.ShouldBindJSON(&clubApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	err = club.ClubApp.ClubCreate(clubApp)
	if err != nil {
		util.GResponseErr(c, CREATEFAILED, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Created")
	util.GResponse(c, CREATED, clubApp)
}


func (club *Club)GGetClubs(c *gin.Context)  {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	clubApp, err := club.ClubApp.GetClubs()
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Get clubs")
	util.GResponse(c, STATUSOK, clubApp)
}

func (club *Club)GGetClub(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		logger.Error(errors.New("Cloud not get the param").Error())
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		return
	}
	clubApp := application.Club{
		Id: id,
	}
	clubApp, err := club.ClubApp.GetClub(clubApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
		return
	}
	logger.Info("Get club")
	util.GResponse(c, STATUSOK, clubApp)
}


func (club *Club)GUpdateClub(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		logger.Error(errors.New("Cloud not get the param").Error())
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		return
	}
	clubApp := application.Club{}
	err := c.ShouldBindJSON(&clubApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
		return
	}
	clubApp.Id = id
	resultUpdate, err := club.ClubApp.Update(clubApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
		return
	}
	logger.Info("Updated")
	util.GResponse(c, STATUSOK, resultUpdate)
}
