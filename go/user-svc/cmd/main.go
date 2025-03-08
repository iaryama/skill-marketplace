package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"skill-marketplace/user-svc/db"
	userGrpc "skill-marketplace/user-svc/grpc"
	"skill-marketplace/user-svc/handlers"
	"skill-marketplace/user-svc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"skill-marketplace/user-svc/config"
)

func startRESTServer(wg *sync.WaitGroup) {
	defer wg.Done()

	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	r.POST("/providers", handlers.CreateProvider)

	fmt.Println("REST API running on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}

func startGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	user_proto.RegisterUserServiceServer(grpcServer, &userGrpc.UserServiceServer{})

	fmt.Println("gRPC User Service running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	db.ConnectDatabase()

	var wg sync.WaitGroup
	wg.Add(2)

	go startRESTServer(&wg)
	go startGRPCServer(&wg)

	wg.Wait()
}
