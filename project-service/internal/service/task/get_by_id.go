package task

import (
	"context"
	"project-service/internal/errorz"
	"project-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.Task, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errorz.ErrUserIDNotSet
	}

	task, err := s.taskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	project, err := s.projectRepository.GetByID(ctx, task.ProjectId)
	if err != nil {
		return nil, err
	}

	for _, pUserID := range project.Users {
		if userID == pUserID {
			return s.taskRepository.GetByID(ctx, id)
		}
	}

	return nil, errorz.ErrTaskAccessForbidden
}
