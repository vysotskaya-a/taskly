package task

import (
	"context"
	"project-service/internal/errorz"
	"project-service/internal/models"
)

func (s *Service) Create(ctx context.Context, task *models.Task) (string, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return "", errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, task.ProjectId)
	if err != nil {
		return "", err
	}

	for _, pUserID := range project.Users {
		if userID == pUserID {
			return s.taskRepository.Create(ctx, task)
		}
	}

	return "", errorz.ErrTaskAccessForbidden
}
