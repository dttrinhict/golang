package handler

import (
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	"net/http"
)

func (user *User)GUserCreate(c *gin.Context) {
	userApp := application.User{}
	err := c.ShouldBindJSON(&userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	err = user.UserApp.UserCreate(userApp)
	if err != nil {
		util.GResponseErr(c, http.StatusNotFound, err.Error())
	}
	util.GResponse(c, http.StatusCreated, userApp)
}


func (user *User)GGetUsers(c *gin.Context)  {
	userApp, err := user.UserApp.GetUsers()
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	util.GResponse(c, http.StatusOK, userApp)
}