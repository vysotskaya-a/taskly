package user

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "user-service/pkg/api/user_v1"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Логика получения пользователя
	user, err := s.userService.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Grade:     user.Grade,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}
