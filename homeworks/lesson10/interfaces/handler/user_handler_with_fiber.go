package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
)

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

func (user *User)GetUsers(c *fiber.Ctx) (err error)  {
	userApp, err := user.UserApp.GetUsers()
	if err != nil {
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(userApp)
}

func (user *User)GetUser(c *fiber.Ctx) (err error)  {
	id := c.Params("id")
	userApp := application.User{
		Id: id,
	}
	userApp, err = user.UserApp.GetUser(userApp)
	if err != nil {
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(userApp)
}


func (user *User)UpdateUser(c *fiber.Ctx) (err error)  {
	userApp := application.User{}
	err = c.BodyParser(&userApp)
	userApp, err = user.UserApp.Update(userApp)
	if err != nil {
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(userApp)
}
