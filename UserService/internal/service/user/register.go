package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"user-service/internal/models"
)

func (s *Service) Register(ctx context.Context, user *models.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	return s.userRepository.Create(ctx, user)
}
