package entities

type User struct { //Tên Entity --> biên thành snake_case trong Postgresql
	tableName  struct{} `pg:"golang.users"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk" json:"id"`      //pk--> đây là primary key
	Name       string   `pg:"name" json:"name"`//-> name, kiểu string --> text
	Email      string	`pg:"email" json:"email"`
	Mobile     string	`pg:"mobile" json:"mobile"`
	Int_roles  []int    `pg:"int_roles,array" json:"int_roles"`  //Quy ước IntRoles --> int_roles snake case
	Enum_roles []string `pg:"enum_roles,array" json:"enum_roles"` //kiểu cột là array lưu string
}