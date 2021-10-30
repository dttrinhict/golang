package model

import "golang/homeworks/lesson10/entities"

type Club struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Users []*entities.User `json:"users"`
	Members []*entities.Member `json:"members"`
}
