package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"golang/homeworks/lesson10/interfaces/handler"
)

func NewRouter(user *handler.User) *fiber.App {
	app := fiber.New()
	app.Post("/users", user.UserCreate)
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	return app
}
