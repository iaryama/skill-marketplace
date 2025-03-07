package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"skill-marketplace/user-svc/models"
	"skill-marketplace/user-svc/utils/constants"
)

var DB *gorm.DB
var (
	DBHost     = constants.DBHost
	DBUser     = constants.DBUser
	DBPassword = constants.DBPassword
	DBName     = constants.DBName
	DBPort     = constants.DBPort
)

func ConnectDatabase() {
	// Construct the DSN using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Provider{})
	DB = database
}
