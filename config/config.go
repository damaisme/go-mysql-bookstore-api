package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

func SetupDatabaseConn() *gorm.DB {
	dsn := "root:my-secret-pw@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect to database")
	}
	return db

}

func CloseDatabaseConn(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close database connection: ", err)
	}
	conn.Close()
}
