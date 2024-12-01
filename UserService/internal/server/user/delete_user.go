package user

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-service/pkg/api/user_v1"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	// Логика удаления пользователя
	err := s.userService.DeleteUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
