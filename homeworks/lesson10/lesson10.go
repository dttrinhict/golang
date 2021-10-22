package main

import (
	"golang/homeworks/lesson10/application"
	domainservices "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/infrastucture/postgress"
	"golang/homeworks/lesson10/infrastucture/repo"
	"golang/homeworks/lesson10/interfaces"
	"golang/homeworks/lesson10/interfaces/handler"
)

func main()  {
	db := postgress.GetIntance()
	user_Postgress_Repo := repo.User_Postgress_Repo(&db)
	domain_User := domainservices.DomainUser(user_Postgress_Repo)
	user_App := application.User_App(domain_User)
	user := handler.NewUser(user_App)
	app := interfaces.NewRouter(user)
	app.Listen(":8080")
}
