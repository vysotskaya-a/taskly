package auth

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "user-service/pkg/api/auth_v1"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	refreshToken, err := s.authService.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("error while logging in")
		return nil, status.Error(codes.Unauthenticated, "Failed to login")
	}

	return &pb.LoginResponse{RefreshToken: refreshToken}, nil
}
