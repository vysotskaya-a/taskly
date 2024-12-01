package user

import (
	"context"
	"user-service/internal/models"
)

func (s *Service) UpdateUser(ctx context.Context, user *models.User) error {
	return s.userRepository.Update(ctx, user)
}
