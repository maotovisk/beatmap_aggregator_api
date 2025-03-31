package database

import (
	"log"
	"simple_api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	log.Println("Starting database...")
	// Database
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
	}
	log.Println("Database connected successfully.")

	db.AutoMigrate(model.ImportModels()...)
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Println("Database not initialized, initializing now...")
		InitDatabase()
	}

	return db
}
