package main

import (
	"fmt"
	"golang/homeworks/lesson10/application"
	domainservices "golang/homeworks/lesson10/domain/services"
	"golang/homeworks/lesson10/infrastucture/databases"
	"golang/homeworks/lesson10/infrastucture/repo"
	"golang/homeworks/lesson10/interfaces"
	"golang/homeworks/lesson10/interfaces/handler"
	"golang/homeworks/lesson10/util/configs"
	"log"
)

func main()  {

	config, err := configs.LoadConfig("./envconfig/")
	if err != nil {
		log.Printf(err.Error())
	}

	pgdb := databases.PgDBIntance()
	mysqldb := databases.MySQLDBIntance(config)

	user_Postgress_Repo := repo.User_Postgress_Repo(pgdb)
	//domain_User := domainservices.DomainUser(user_Postgress_Repo) // user postgress to store data
	user_mysql_repo := repo.User_MySQL_Repo(mysqldb)
	club_mysql_repo := repo.Club_MySQL_Repo(mysqldb)
	user_role_mysql := repo.User_Role_MySQL_Repo(mysqldb)
	user_role_postgress := repo.User_Role_Postgress_Repo(pgdb)
	member_mysql_repo := repo.MemberMySQLRepo(mysqldb)
	fmt.Println(user_Postgress_Repo)
	fmt.Println(user_mysql_repo)
	fmt.Println(user_role_mysql)
	domain_User := domainservices.DomainUser(user_Postgress_Repo) // user "user_Postgress_Repo" for postgress and "user_mysql_repo" for mysql to store data
	user_App := application.User_App(domain_User)
	user := handler.NewUser(user_App)

	domain_club := domainservices.DomainClub(club_mysql_repo) // user mysql to store data
	club_App := application.Club_App(domain_club)
	club := handler.NewClub(club_App)

	domain_user_role := domainservices.DomainUserClub(user_role_postgress) // user "user_club_mysql" for mysql and "" for postgres to store data
	user_club_App := application.User_Club_App(domain_user_role)
	user_club := handler.NewUserClub(user_club_App)

	domainMember := domainservices.DomainMember(member_mysql_repo) // user mysql to store data
	memberApp := application.Member_App(domainMember)
	member := handler.NewMember(memberApp)


	// Run with fiber web engine
	app := interfaces.NewRouter(user, club, user_club, member)
	app.Listen(config.ServerAddress)

	// Run with gin web engine
	//app := interfaces.NewGinServer(user, club, user_club,member)
	//app.Run(config.ServerAddress)

}
