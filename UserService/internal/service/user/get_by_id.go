package user

import (
	"context"
	"user-service/internal/models"
)

func (s *Service) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepository.GetByID(ctx, id)
}
