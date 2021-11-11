package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson12/application"
	"golang/homeworks/lesson12/util"
	ilogger "golang/homeworks/lesson12/util/logger"
)

func (memberClub *MemberClub)GAssignMembersToClub(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		logger.Error(errors.New("Cloud not get the param").Error())
		return
	}
	clubApp := application.Club{
		Id: id,
	}
	membersApp := []application.Member{}
	err := c.ShouldBindJSON(&membersApp)
	membersApp, err = memberClub.MemberClubApp.AssignMembersToClub(clubApp, membersApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Updated")
	util.GResponse(c, STATUSOK, membersApp)
}

func (memberClub *MemberClub)GGetMembersOfClub(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		logger.Error(errors.New("Cloud not get the param").Error())
		return
	}
	clubApp := application.Club{
		Id: id,
	}
	membersApp, err := memberClub.MemberClubApp.GetMembersOfClub(clubApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Updated")
	util.GResponse(c, STATUSOK, membersApp)
}