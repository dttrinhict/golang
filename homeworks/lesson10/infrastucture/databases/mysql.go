package databases

import (
	"errors"
	"golang/homeworks/lesson10/util/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type MySQLDB struct {
	DB *gorm.DB
	Random *rand.Rand
}

var mysqlInstance *MySQLDB

func MySQLDBIntance(config *configs.Config) *MySQLDB  {
	if mysqlInstance == nil {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: config.DBSource, // data source name
			DefaultStringSize: 256, // default size for string fields
			DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		if db == nil {
			err = errors.New("Gorm return nil pointer")
			panic(err)
		}else{
			sqlDB, err := db.DB()
			if err != nil {
				panic(err)
			}
			sqlDB.SetMaxIdleConns(100)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(10*time.Minute)
		}
		//Khởi động engine sinh số ngẫu nhiên
		s1 := rand.NewSource(time.Now().UnixNano())
		random = rand.New(s1)
		mysqlInstance = &MySQLDB{
			DB: db,
			Random: random,
		}
	}
	return mysqlInstance
}