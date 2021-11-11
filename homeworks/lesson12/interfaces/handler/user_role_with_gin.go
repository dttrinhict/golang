package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson12/application"
	"golang/homeworks/lesson12/util"
	"net/http"
)

func (userRole *UserRole)GAssignUserToRole(c *gin.Context) {
	roleApp := application.Role{}
	err := c.ShouldBindJSON(&roleApp)
	if err != nil {
		util.GResponseErr(c, http.StatusBadRequest, err.Error())
	}
	users, err := userRole.UserRoleApp.AssignUserToRole(roleApp)
	if err != nil {
		util.GResponseErr(c, http.StatusNotFound, err.Error())
	}
	util.GResponse(c, http.StatusCreated, users)
}


func (userRole *UserRole)GGetUsersOfRole(c *gin.Context) {
	clubID := c.Param("id")
	if clubID == "" {
		util.GResponseErr(c, http.StatusBadRequest, errors.New("Role id is null").Error())
	}
	roleApp := application.Role{
		Id: clubID,
	}
	users, err := userRole.UserRoleApp.GetUsersOfRole(roleApp)
	if err != nil {
		util.GResponseErr(c, http.StatusNotFound, err.Error())
	}
	util.GResponse(c, http.StatusCreated, users)
}

