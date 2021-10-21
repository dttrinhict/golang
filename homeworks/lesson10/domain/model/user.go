package model

type User struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string	`json:"email"`
	Mobile     string	`json:"mobile"`
	Int_roles  []int    `json:"int_roles"`
	Enum_roles []string `json:"enum_roles"`
}