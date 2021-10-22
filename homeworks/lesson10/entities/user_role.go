package entities



type User_Role struct {
	tableName struct{} `pg:"golang.user_role"`
	User_id   string `pg:"user_id" gorm:"user_id" json:"user_id"`
	Role_id   int	`pg:"role_id" gorm:"role_id" json:"role_id"`
	Role Role
	User User
}

func (user_role *User_Role) TableName() string  {
	return "user_role"
}