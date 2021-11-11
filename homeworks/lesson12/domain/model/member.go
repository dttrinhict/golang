package model

import "golang/homeworks/lesson12/entities"

type Member struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string	`json:"email"`
	Mobile     string	`json:"mobile"`
	Clubs []entities.Club `json:"clubs"`
}