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
	user_proto.UnimplementedUserServiceServer
}

// Get User by ID
func (s *UserServiceServer) GetUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.UserResponse, error) {
	var user models.User
	if err := db.DB.First(&user, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		log.Println("Database error:", err)
		return nil, err
	}

	return &user_proto.UserResponse{
		Id:    string(rune(user_proto.ID)),
		Name:  user_proto.Name,
		Email: user_proto.Email,
	}, nil
}

// Get Provider by ID
func (s *UserServiceServer) GetProvider(ctx context.Context, req *user_proto.GetProviderRequest) (*user_proto.ProviderResponse, error) {
	var provider models.Provider
	if err := db.DB.First(&provider, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		log.Println("Database error:", err)
		return nil, err
	}

	return &user_proto.ProviderResponse{
		Id:    string(rune(provider.ID)),
		Type:  string(provider.Type),
		Name:  provider.Name,
		Email: provider.Email,
	}, nil
}
