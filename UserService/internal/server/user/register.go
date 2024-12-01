package user

import (
	"context"
	"user-service/internal/models"
	pb "user-service/pkg/api/user_v1"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Логика регистрации пользователя
	user := &models.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Grade:    req.GetGrade(),
	}

	userID, err := s.userService.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Id: userID,
	}, nil
}
