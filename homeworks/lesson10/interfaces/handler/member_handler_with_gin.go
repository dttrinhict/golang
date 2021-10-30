package handler

import (
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	ilogger "golang/homeworks/lesson10/util/logger"
	"net/http"
)

func (member *Member)GMemberCreate(c *gin.Context) {
	logger := ilogger.NewLogger(ilogger.FactoryLogger{GinContext: c})
	memberApp := application.Member{}
	err := c.ShouldBindJSON(&memberApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
		logger.Log(err)
		return
	}
	err = member.memberApp.Create(memberApp)
	if err != nil {
		//util.GResponseErr(c, http.StatusNotFound, err.Error())
		logger.Log(err)
		return
	}
	logger.Info("Created", c.Request.RequestURI)
	util.GResponse(c, http.StatusCreated, memberApp)
}


func (member *Member)GGetMembers(c *gin.Context)  {
	memberApp, err := member.memberApp.GetMembers()
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, memberApp)
}

func (member *Member)GGetMember(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {

	}
	memberApp := application.Member{
		Id: id,
	}
	memberApp, err := member.memberApp.GetMember(memberApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, memberApp)
}


func (member *Member)GUpdateMember(c *gin.Context) {
	memberApp := application.Member{}
	err := c.ShouldBindJSON(&memberApp)
	memberApp, err = member.memberApp.Update(memberApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, memberApp)
}