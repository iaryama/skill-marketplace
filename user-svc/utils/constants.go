package constants

import (
	"os"
)

// Exported database configuration constants
var (
	DBHost     string = os.Getenv("DB_HOST")
	DBPort     string = os.Getenv("DB_PORT")
	DBUser     string = os.Getenv("DB_USER")
	DBPassword string = os.Getenv("DB_PASSWORD")
	DBName     string = os.Getenv("DB_NAME")
)
