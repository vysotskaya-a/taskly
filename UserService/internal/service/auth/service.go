package auth

import (
	"user-service/internal/config"
	"user-service/internal/repository"
)

type Service struct {
	jwtConfig config.JWTConfig

	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}
