package project

import (
	"context"
	"project-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.Project, error) {
	return s.projectRepository.GetByID(ctx, id)
}
