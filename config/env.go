package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Environment variables loaded successfully")
}

type DatabaseConfig struct {
	url     string
}
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url:     os.Getenv("DATABASE_URL"),
	}
}