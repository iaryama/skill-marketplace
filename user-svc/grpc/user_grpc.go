package grpc

import (
	"context"
	"gorm.io/gorm"
	"log"
	"skill-marketplace/user-svc/db"
	"skill-marketplace/user-svc/models"
	"skill-marketplace/user-svc/proto"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
}

// Get User by ID
func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	var user models.User
	if err := db.DB.First(&user, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		log.Println("Database error:", err)
		return nil, err
	}

	return &proto.UserResponse{
		Id:    string(rune(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Get Provider by ID
func (s *UserServiceServer) GetProvider(ctx context.Context, req *proto.GetProviderRequest) (*proto.ProviderResponse, error) {
	var provider models.Provider
	if err := db.DB.First(&provider, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		log.Println("Database error:", err)
		return nil, err
	}

	return &proto.ProviderResponse{
		Id:    string(rune(provider.ID)),
		Type:  string(provider.Type),
		Name:  provider.Name,
		Email: provider.Email,
	}, nil
}
