package entities

type Club struct {
	tableName  struct{} `pg:"golang.clubs"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk" gorm:"Id,primaryKey" json:"id"`      //pk--> đây là primary key
	Name       string   `pg:"name" gorm:"name" json:"name"`//-> name, kiểu string --> text
	Users	[]*User `pg:"many2many:user_club" gorm:"many2many:user_club" json:"users"`
	Members	[]*Member `pg:"many2many:member_club" gorm:"many2many:member_club" json:"members"`
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

type Member_Club struct {
	tableName struct{} `pg:"golang.member_club"`
	Member_id   string `pg:"member_id" gorm:"member_id" json:"member_id"`
	Club_id   string	`pg:"club_id" gorm:"club_id" json:"club_id"`
}

func (member_club *Member_Club) TableName() string  {
	return "member_club"
}