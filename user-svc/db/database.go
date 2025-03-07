package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"skill-marketplace/user-svc/utils/constants"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Construct the DSN using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		constants.DB_HOST, constants.DB_USERNAME, constants.DB_PASSWORD, constants.DB_NAME, constants.DB_PORT,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Provider{})
	DB = database
}
