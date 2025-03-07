package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
	DatabaseURL string
	Port        int
}

var AppConfig Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	AppConfig.DBHost = os.Getenv("DB_HOST")
	AppConfig.DBUser = os.Getenv("DB_USER")
	AppConfig.DBPassword = os.Getenv("DB_PASSWORD")
	AppConfig.DBName = os.Getenv("DB_NAME")
	AppConfig.DatabaseURL = os.Getenv("DATABASE_URL")

	portStr := os.Getenv("DB_PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("invalid DB_PORT: %w", err)
		}
		AppConfig.DBPort = port
	}

	portStr = os.Getenv("PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("invalid PORT: %w", err)
		}
		AppConfig.Port = port
	}

	return nil
}
