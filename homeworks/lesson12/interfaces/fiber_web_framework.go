package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"golang/homeworks/lesson12/interfaces/handler"
)

func NewRouter(user *handler.User, club *handler.Club, userRole *handler.UserRole, member *handler.Member) *fiber.App {
	app := fiber.New()
	userGroup := app.Group("/user/")
	{
		userGroup.Post("/create", user.UserCreate)
		userGroup.Get("/get-users", user.GetUsers)
		userGroup.Get("/get-user/:id", user.GetUser)
		userGroup.Put("/update-user/:id", user.UpdateUser)
		userGroup.Delete("/delete-user/:id", user.DeleteUser)
	}
	return app
}
