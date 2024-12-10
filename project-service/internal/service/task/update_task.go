package task

import (
	"context"
	"project-service/internal/models"
)

func (s *ServiceTask) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.UpdateTask(ctx, task)
}
