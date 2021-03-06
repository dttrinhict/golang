package entities

type User struct { //Tên Entity --> biên thành snake_case trong Postgresql
	tableName  struct{} `pg:"golang.users"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`      //pk--> đây là primary key
	Name       string   `pg:"name" gorm:"name" json:"name"`//-> name, kiểu string --> text
	Email      string	`pg:"email" gorm:"email" json:"email"`
	Mobile     string	`pg:"mobile" gorm:"mobile" json:"mobile"`
	Roles []Role `pg:"many2many:golang.user_role" gorm:"many2many:user_role" json:"roles"`
}

func (user *User) TableName() string  {
	return "users"
}