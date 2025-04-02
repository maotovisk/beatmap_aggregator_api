package database

import (
	"beatmap_aggregator_api/config"
	"beatmap_aggregator_api/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	config := config.GetConfig()
	var err error
	log.Println("Starting database...")
	// Database
	db, err = gorm.Open(sqlite.Open(config.Sqlite.Path), &gorm.Config{})
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
