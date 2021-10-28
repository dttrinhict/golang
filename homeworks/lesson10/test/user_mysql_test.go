package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang/homeworks/lesson10/entities"
	"golang/homeworks/lesson10/infrastucture/databases"
	"golang/homeworks/lesson10/infrastucture/repo"
	"testing"
)

func TestU(t *testing.T) {
	
}

func Test_GetClub_MySQL(t *testing.T) {
	mySQL := databases.MySQLDBIntance()
	ClubMySQLRepo := repo.Club_MySQL_Repo(mySQL)
	member, err := ClubMySQLRepo.GetClub(entities.Club{Id: "Rwd3JJHH"})
	assert.Nil(t, err)

	fmt.Println(member)

	fmt.Println(member.Name)
	for _, c := range member.User {
		fmt.Println("	" + c.Name)
	}
}
