package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson12/application"
	"golang/homeworks/lesson12/util"
	ilogger "golang/homeworks/lesson12/util/logger"
)

func (member *Member)GMemberCreate(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	memberApp := application.Member{}
	err := c.ShouldBindJSON(&memberApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
		return
	}
	err = member.memberApp.Create(memberApp)
	if err != nil {
		util.GResponseErr(c, CREATEFAILED, err.Error())
		logger.Error(err.Error())
		return
	}
	logger.Info("Created")
	util.GResponse(c, CREATED, memberApp)
}


func (member *Member)GGetMembers(c *gin.Context)  {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	memberApp, err := member.memberApp.GetMembers()
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
		return
	}
	logger.Info("Get memebers")
	util.GResponse(c, STATUSOK, memberApp)
}

func (member *Member)GGetMember(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		logger.Error(errors.New("Cloud not get the param").Error())
		return
	}
	memberApp := application.Member{
		Id: id,
	}
	memberApp, err := member.memberApp.GetMember(memberApp)
	if err != nil {
		logger.Error(err.Error())
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		return
	}
	logger.Info("Get member")
	util.GResponse(c, STATUSOK, memberApp)
}


func (member *Member)GUpdateMember(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		logger.Error(errors.New("Cloud not get the param").Error())
		return
	}
	memberApp := application.Member{}
	memberApp.Id = id
	err := c.ShouldBindJSON(&memberApp)
	memberApp, err = member.memberApp.Update(memberApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Updated")
	util.GResponse(c, STATUSOK, memberApp)
}

func (member *Member) GDeleteMember(c *gin.Context) {
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	id, ok := c.Params.Get("id")
	if !ok {
		util.GResponseErr(c, STATUSBEDREQUEST, errors.New("Cloud not get the param").Error())
		logger.Error(errors.New("Cloud not get the param").Error())
		return
	}
	memberApp := application.Member{}
	memberApp.Id = id
	err := c.ShouldBindJSON(&memberApp)
	memberApp, err = member.memberApp.Update(memberApp)
	if err != nil {
		util.GResponseErr(c, STATUSBEDREQUEST, err.Error())
		logger.Error(err.Error())
	}
	logger.Info("Updated")
	util.GResponse(c, STATUSOK, memberApp)
}