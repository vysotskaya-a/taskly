package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "user-service/pkg/api/auth_v1"
)

func (s *Server) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error) {
	refreshToken, err := s.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &pb.GetRefreshTokenResponse{RefreshToken: refreshToken}, nil
}
