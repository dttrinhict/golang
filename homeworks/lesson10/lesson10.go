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
	user_RepoPostgress_Create := repo.User_RepoPostgress_Create(&db)
	user_Repo_Postgress_Read := repo.User_RepoPostgress_Read(&db)
	domain_User := domainservices.DomainUser(user_RepoPostgress_Create, user_Repo_Postgress_Read)
	user_App := application.User_App(domain_User)
	user := handler.NewUser(user_App)
	app := interfaces.NewRouter(user)
	app.Listen(":8080")
}
