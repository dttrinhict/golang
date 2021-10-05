package main

import (
	"github.com/dttrinhict/golang/lesson07/homeworks/models"
)

func main()  {
	user := models.NewUser()
	server := models.NewServer(user)
	models.NewHttp(server)
}
