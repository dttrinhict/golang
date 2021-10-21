package entities



type User_Role struct {
	tableName struct{} `pg:"golang.user_role"`
	User_id   string `json:"user_id"`
	Role_id   int	`json:"role_id"`
}
