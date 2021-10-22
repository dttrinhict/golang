package databases

import (
	"gorm.io/gorm"
	"math/rand"
)

type MySQLDB struct {
	DB *gorm.DB
	Random *rand.Rand
}