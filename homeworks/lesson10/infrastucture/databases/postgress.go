package databases

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"golang/homeworks/lesson10/entities"
	"math/rand"
	"time"
)

var (
	//DB     interface{}    //kết nối vào CSDL Postgresql
	random *rand.Rand // Đối tượng dùng để tạo random number
)

var PgDB *PostgressDB

type PostgressDB struct {
	DB *pg.DB
	Random *rand.Rand
}

func PgDBIntance() *PostgressDB {
	if PgDB == nil {
		DB := pg.Connect(&pg.Options{
			Addr:     "localhost:5432",
			User:     "postgres",
			Password: "abc123",
			//Database: "golang",
		})
		//Đăng ký bảng quan hệ nhiều - nhiều
		orm.RegisterTable(&entities.User_Role{})
		//if err := DB.Ping(context.Background()); err != nil {
		//	panic(err)
		//}
		//Log các câu lệnh SQL thực thi để debug
		DB.AddQueryHook(dbLogger{}) //Log query to console

		//Khởi động engine sinh số ngẫu nhiên
		s1 := rand.NewSource(time.Now().UnixNano())
		random = rand.New(s1)

		PgDB = &PostgressDB{
			DB:     DB,
			Random: random,
		}
	}
	return PgDB
}

type dbLogger struct{}

// Hàm hook (móc câu vào lệnh truy vấn) để in ra câu lệnh SQL query
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// Hàm hook chạy sau khi query được thực thi
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println(string(bytes))
	return nil
}
