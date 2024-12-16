package auth

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "user-service/pkg/api/auth_v1"
)

func (s *Server) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error) {
	accessToken, err := s.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		log.Error().Err(err).Msg("error while getting access token")
		return nil, status.Error(codes.Unauthenticated, "Failed to get access token.")
	}

	return &pb.GetAccessTokenResponse{AccessToken: accessToken}, nil
}
