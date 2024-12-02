package service

import (
	"context"
	"user-service/internal/models"
	pb "user-service/pkg/api/auth_v1"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) (string, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type AuthService interface {
	Login(ctx context.Context, req *pb.LoginRequest) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
