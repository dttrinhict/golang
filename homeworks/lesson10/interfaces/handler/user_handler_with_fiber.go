package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang/homeworks/lesson10/application"
	"golang/homeworks/lesson10/util"
	"golang/homeworks/lesson10/util/logger"
	"net/http"
)

/*
Lấy danh sách user
*/
func (user *User)UserCreate(c *fiber.Ctx) (err error) {
	log := logger.NewFactoryZapLogger(logger.FactoryLogger{FiberContext: c})
	userApp := application.User{}
	err = c.BodyParser(&userApp)
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	userResult, err := user.UserApp.UserCreate(userApp)
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusNotFound, err.Error())
	}
	log.Info("Create user successful")
	return util.FResponse(c, http.StatusCreated, userResult)
}

func (user *User)GetUsers(c *fiber.Ctx) (err error)  {
	log := logger.NewFactoryZapLogger(logger.FactoryLogger{FiberContext: c})
	userApp, err := user.UserApp.GetUsers()
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	log.Info("Get users successful")
	return util.FResponse(c, http.StatusOK, userApp)
}

func (user *User)GetUser(c *fiber.Ctx) (err error)  {
	log := logger.NewFactoryZapLogger(logger.FactoryLogger{FiberContext: c})
	id := c.Params("id")
	userApp := application.User{
		Id: id,
	}
	userApp, err = user.UserApp.GetUser(userApp)
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	log.Info("Get user "+ userApp.Id)
	return util.FResponse(c, http.StatusOK, userApp)
}


func (user *User)UpdateUser(c *fiber.Ctx) (err error)  {
	log := logger.NewFactoryZapLogger(logger.FactoryLogger{FiberContext: c})
	userRequestBody := application.User{}
	err = c.BodyParser(&userRequestBody)
	id := c.Params("id")
	userApp := application.User{
		Id: id,
	}
	userApp, err = user.UserApp.GetUser(userApp)
	if err != nil || userApp.Id == "" {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	userRequestBody.Id = id
	userApp, err = user.UserApp.Update(userRequestBody)
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	log.Info("Updating the user "+ userApp.Id + "is successful")
	return c.JSON(userApp)
}


func (user *User)UDeleteUser(c *fiber.Ctx) (err error)  {
	log := logger.NewFactoryZapLogger(logger.FactoryLogger{FiberContext: c})
	id := c.Params("id")
	userApp := application.User{
		Id: id,
	}
	userApp, err = user.UserApp.GetUser(userApp)
	if err != nil || userApp.Id == "" {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	users, err := user.UserApp.DeleteUser(userApp)
	if err != nil {
		log.Error(err.Error())
		return util.ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}
	log.Info("Deleting the user "+ userApp.Id + "is successful")
	return c.JSON(users)
}
