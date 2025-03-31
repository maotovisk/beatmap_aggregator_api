package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Osu    OsuConfig
	Sqlite SqliteConfig
	Web    WebConfig
}

var config AppConfig

func loadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = env.Parse(&config)
	config, err = env.ParseAs[AppConfig]()
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() AppConfig {
	err := loadConfig()
	if err != nil {
		log.Println("config: ", config)
	}

	return config
}
