package entities

type Role struct {
	tableName  struct{} `pg:"golang.role"`
	Id int `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`
	Name string `pg:"name" gorm:"name" json:"name"`
	User_Role []User_Role
}

func (role *Role) TableName() string  {
	return "role"
}