package user

import (
	"context"
	"google.golang.org/grpc/metadata"
	"user-service/internal/errorz"
	"user-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.User, error) {
	var userID string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userID = md["user_id"][0]
	}
	if !ok {
		return nil, errorz.ErrUserIDNotSet
	}

	if userID != id {
		return nil, errorz.ErrUserAccessDenied
	}

	return s.userRepository.GetByID(ctx, id)
}
