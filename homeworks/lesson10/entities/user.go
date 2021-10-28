package entities

type User struct { //Tên Entity --> biên thành snake_case trong Postgresql
	tableName  struct{} `pg:"golang.users"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`      //pk--> đây là primary key
	Name       string   `pg:"name" gorm:"name" json:"name"`//-> name, kiểu string --> text
	Email      string	`pg:"email" gorm:"email" json:"email"`
	Mobile     string	`pg:"mobile" gorm:"mobile" json:"mobile"`
	Clubs	[]*Club `pg:"many2many:user_club" gorm:"many2many:user_club" json:"club"`
	User_Role []User_Role
}

func (user *User) TableName() string  {
	return "users"
}