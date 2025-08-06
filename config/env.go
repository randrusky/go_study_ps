package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Environment variables loaded successfully")
}

func getString(rey, defaultValue string) string {
	value := os.Getenv(rey)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

type DatabaseConfig struct {
	url     string
}
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url:    getString("DATABASE_URL", ""),
	}
}