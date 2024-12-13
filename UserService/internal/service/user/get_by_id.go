package user

import (
	"context"
	"user-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepository.GetByID(ctx, id)
}
