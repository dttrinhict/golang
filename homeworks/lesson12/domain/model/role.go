package model

import "golang/homeworks/lesson12/entities"

type Role struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Users []entities.User `json:"users"`
}
