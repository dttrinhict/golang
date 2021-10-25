package main

import (
	"fmt"
	"golang/homeworks/lesson10/application"
	domainservices "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/infrastucture/databases"
	"golang/homeworks/lesson10/infrastucture/repo"
	"golang/homeworks/lesson10/interfaces"
	"golang/homeworks/lesson10/interfaces/handler"
)

func main()  {
	pgdb := databases.PgDBIntance()
	mysqldb := databases.MySQLDBIntance()

	user_Postgress_Repo := repo.User_Postgress_Repo(pgdb)
	//domain_User := domainservices.DomainUser(user_Postgress_Repo) // user postgress to store data
	user_mysql_repo := repo.User_MySQL_Repo(mysqldb)
	fmt.Println(user_Postgress_Repo)
	domain_User := domainservices.DomainUser(user_mysql_repo) // user mysql to store data
	user_App := application.User_App(domain_User)
	user := handler.NewUser(user_App)
	// Run with fiber web engine
	//app := interfaces.NewRouter(user)
	//app.Listen(":8080")

	// Run with gin web engine
	app := interfaces.NewGinServer(user)
	app.Run(":8080")

}
