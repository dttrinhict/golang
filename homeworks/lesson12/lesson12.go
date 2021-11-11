package main

import (
	"fmt"
	"golang/homeworks/lesson12/application"
	domainservices "golang/homeworks/lesson12/domain/services"
	"golang/homeworks/lesson12/infrastucture/databases"
	"golang/homeworks/lesson12/infrastucture/repo"
	"golang/homeworks/lesson12/interfaces"
	"golang/homeworks/lesson12/interfaces/handler"
	"golang/homeworks/lesson12/util/configs"
	"log"
)

func main()  {

	config, err := configs.LoadConfig("./homeworks/lesson12/envconfig/")
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
	member_club_mysql := repo.MemeberClubMySQLRepo(mysqldb)
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


	domain_Member_Club := domainservices.DomainMemberClub(member_club_mysql) // user mysql to store data
	member_club_App := application.Member_Club_App(domain_Member_Club)
	member_club := handler.NewMemberClub(member_club_App)


	// Run with fiber web engine
	//app := interfaces.NewRouter(user, club, user_club, member)
	//app.Listen(config.ServerAddress)

	// Run with gin web engine
	app := interfaces.NewGinServer(user, club, user_club, member, member_club)
	app.Run(config.ServerAddress)

}
