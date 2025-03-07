package main

import (
	"github.com/gin-gonic/gin"
	"skill-marketplace/user-svc/db"
	"skill-marketplace/user-svc/handlers"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.POST("/providers", handlers.CreateProvider)

	r.Run(":8081")
}
