package configs

import (
	"log"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func ConnectDB() *gorm.DB{
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlserver.Open(GetEnv("DB_DSN")))
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}
		log.Println("Connected to database")
	})

	return db
}