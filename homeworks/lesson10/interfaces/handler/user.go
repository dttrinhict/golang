package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
)

type User struct {
	UserApp application.UserApp
}

func NewUser(userApp application.UserApp) *User {
	return &User{
		UserApp: userApp,
	}
}

/*
Lấy danh sách user
*/
func (user *User)UserCreate(c *fiber.Ctx) (err error) {
	userApp := application.User{}
	err = c.BodyParser(&userApp)
	if err != nil {
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	err = user.UserApp.UserCreate(userApp)
	if err != nil {
		return util.ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(userApp)
}

func (user *User)GetUser(c *fiber.Ctx) (err error)  {
	userApp, err := user.UserApp.GetUsers()
	if err != nil {
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(userApp)
}
