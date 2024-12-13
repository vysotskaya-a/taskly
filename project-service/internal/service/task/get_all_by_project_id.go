package task

import (
	"context"
	"project-service/internal/errorz"
	"project-service/internal/models"
)

func (s *Service) GetAllByProjectID(ctx context.Context, projectID string) ([]*models.Task, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	for _, pUserID := range project.Users {
		if userID == pUserID {
			return s.taskRepository.GetAllByProjectID(ctx, projectID)
		}
	}

	return nil, errorz.ErrTaskAccessForbidden
}
