package task

import (
	"context"
	"project-service/internal/errorz"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return errorz.ErrUserIDNotSet
	}

	task, err := s.taskRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	project, err := s.projectRepository.GetByID(ctx, task.ProjectId)
	if err != nil {
		return err
	}

	if userID != project.AdminID {
		return errorz.ErrTaskAccessForbidden
	}

	return s.taskRepository.Delete(ctx, id)
}
