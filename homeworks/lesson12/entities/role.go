package entities

type Role struct {
	tableName  struct{} `pg:"golang.role"`
	Id string `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`
	Name string `pg:"name" gorm:"name" json:"name"`
	Users []User `pg:"many2many:golang.user_role" gorm:"many2many:user_role" json:"users"`
}

func (role *Role) TableName() string  {
	return "role"
}