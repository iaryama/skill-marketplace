package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"skill-marketplace/user-svc/config"
	"skill-marketplace/user-svc/models"
)

var DB *gorm.DB

// Connection pool settings
const (
	maxOpenConns    = 25 // Max open connections
	maxIdleConns    = 10 // Max idle connections
	connMaxLifetime = 5  // Connection lifetime in minutes
)

// ConnectDatabase initializes the database with connection pooling and retry plugin
func ConnectDatabase() {
	var err error
	var database *gorm.DB

	// Construct DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	// Open DB connection
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pooling
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute * connMaxLifetime)

	// Register custom GORM plugin for retries
	if err := database.Use(&RetryPlugin{}); err != nil {
		log.Fatalf("Failed to register retry plugin: %v", err)
	}

	// Auto-migrate models
	if err := database.AutoMigrate(&models.User{}, &models.Provider{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	DB = database
	log.Println("Database connection established successfully with pooling and automatic retries.")
}
