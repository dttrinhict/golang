package postgress

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-pg/pg/v10"
)

var (
	DB     *pg.DB     //kết nối vào CSDL Postgresql
	random *rand.Rand // Đối tượng dùng để tạo random number
)

var PgDB PostgressDB

/*
Hàm init này luôn chạy đầu tiên của package repo
*/
func _init() {
	//Mở kết nối vào CSDL Postgresql
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "abc123",
		Database: "golang",
	})

	//Log các câu lệnh SQL thực thi để debug
	DB.AddQueryHook(dbLogger{}) //Log query to console

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

}

type PostgressDB struct {
	DB *pg.DB
	Random *rand.Rand
}

func GetIntance() PostgressDB {
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "abc123",
		//Database: "golang",
	})
	//Log các câu lệnh SQL thực thi để debug
	DB.AddQueryHook(dbLogger{}) //Log query to console

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

	PgDB = PostgressDB{
		DB: DB,
		Random: random,
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
