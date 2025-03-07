package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"skill-marketplace/user-svc/config"
	"skill-marketplace/user-svc/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	var (
		DBHost     = config.AppConfig.DBHost
		DBUser     = config.AppConfig.DBUser
		DBPassword = config.AppConfig.DBPassword
		DBName     = config.AppConfig.DBName
		DBPort     = config.AppConfig.DBPort
	)

	// Construct the DSN using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Provider{})
	DB = database
}
