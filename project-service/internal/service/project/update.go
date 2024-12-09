package project

import (
	"context"
	"project-service/internal/models"
)

func (s *Service) Update(ctx context.Context, project *models.Project) error {
	return s.projectRepository.Update(ctx, project)
}
