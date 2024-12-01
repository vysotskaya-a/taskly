package user

import (
	"user-service/internal/service"
	pb "user-service/pkg/api/user_v1"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewServer(userService service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
