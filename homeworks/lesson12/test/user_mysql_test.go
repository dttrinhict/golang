package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang/homeworks/lesson12/entities"
	"golang/homeworks/lesson12/infrastucture/databases"
	"golang/homeworks/lesson12/infrastucture/repo"
	"golang/homeworks/lesson12/util/configs"
	"log"
	"testing"
)

func TestU(t *testing.T) {
	
}

func Test_GetClub_MySQL(t *testing.T) {
	config, err := configs.LoadConfig("./homeworks/lesson12/envconfig/")
	if err != nil {
		log.Printf(err.Error())
	}
	mySQL := databases.MySQLDBIntance(config)
	ClubMySQLRepo := repo.Club_MySQL_Repo(mySQL)
	member, err := ClubMySQLRepo.GetClub(entities.Club{Id: "Rwd3JJHH"})
	assert.Nil(t, err)

	fmt.Println(member)

	fmt.Println(member.Name)
	for _, c := range member.Members {
		fmt.Println("	" + c.Name)
	}
}
