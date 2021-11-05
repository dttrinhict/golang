package handler

import (
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	ilogger "golang/homeworks/lesson10/util/logger"
	"net/http"
)

func (user *User)GUserCreate(c *gin.Context) {
	//logger := ilogger.NewLogger(ilogger.FactoryLogger{GinContext: c})
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	userApp := application.User{}
	err := c.ShouldBindJSON(&userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
		logger.Error(err.Error())
		return
	}
	err = user.UserApp.UserCreate(userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusConflict, err.Error())
		logger.Error(err.Error())
		return
	}
	logger.Info("Created")
	util.GResponse(c, http.StatusCreated, userApp)
}


func (user *User)GGetUsers(c *gin.Context)  {
	//logger := ilogger.NewLogger(ilogger.FactoryLogger{GinContext: c})
	logger := ilogger.NewFactoryZapLogger(ilogger.FactoryLogger{GinContext: c})
	userApp, err := user.UserApp.GetUsers()
	if err != nil {
		logger.Error(err.Error())
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	logger.Info("Get users")
	util.GResponse(c, http.StatusOK, userApp)
}

func (user *User)GGetUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {

	}
	userApp := application.User{
		Id: id,
	}
	userApp, err := user.UserApp.GetUser(userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}

	util.GResponse(c, http.StatusOK, userApp)
}


func (user *User)GUpdateUser(c *gin.Context) {
	userApp := application.User{}
	err := c.ShouldBindJSON(&userApp)
	userApp, err = user.UserApp.Update(userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, userApp)
}