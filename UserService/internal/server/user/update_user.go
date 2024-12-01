package user

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-service/internal/models"
	pb "user-service/pkg/api/user_v1"
)

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	// Логика обновления пользователя
	user := &models.User{
		ID:       req.GetId(),
		Password: req.GetPassword(),
		Grade:    req.GetGrade(),
	}

	err := s.userService.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
