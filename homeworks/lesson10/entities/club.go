package entities

type Club struct {
	tableName  struct{} `pg:"golang.clubs"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`      //pk--> đây là primary key
	Name       string   `pg:"name" gorm:"name" json:"name"`//-> name, kiểu string --> text
	Users	[]*User `pg:"many2many:user_club" gorm:"many2many:user_club" json:"user"`
}

func (club *Club) TableName() string  {
	return "clubs"
}

type User_Club struct {
	tableName struct{} `pg:"golang.user_club"`
	User_id   string `pg:"user_id" gorm:"user_id" json:"user_id"`
	Club_id   string	`pg:"club_id" gorm:"club_id" json:"club_id"`
}

func (user_club *User_Club) TableName() string  {
	return "user_club"
}