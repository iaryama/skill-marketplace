package grpc

import (
	"context"

	"fmt"
	"gorm.io/gorm"

	"skill-marketplace/user-svc/db"
	"skill-marketplace/user-svc/models"
	"skill-marketplace/user-svc/proto"
)

type UserServiceServer struct {
	user_proto.UnimplementedUserServiceServer
}

// GetUser retrieves a user by ID
func (s *UserServiceServer) GetUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.UserResponse, error) {
	var user models.User
	if err := db.DB.First(&user, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		fmt.Println("Database error:", err)
		return nil, err
	}

	return &user_proto.UserResponse{
		Id:    fmt.Sprintf("%d", user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetProvider retrieves a provider by ID
func (s *UserServiceServer) GetProvider(ctx context.Context, req *user_proto.GetProviderRequest) (*user_proto.ProviderResponse, error) {
	var provider models.Provider
	if err := db.DB.First(&provider, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		fmt.Println("Database error:", err)
		return nil, err
	}

	return &user_proto.ProviderResponse{
		Id:    fmt.Sprintf("%d", provider.ID),
		Type:  string(provider.Type),
		Name:  *provider.CompanyName,
		Email: provider.Email,
	}, nil
}
