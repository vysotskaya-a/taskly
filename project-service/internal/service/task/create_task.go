package task

import (
	"context"
	"project-service/internal/models"
)

func (s *ServiceTask) CreateTask(ctx context.Context, task *models.Task) (string, error) {
	return s.taskRepo.CreateTask(ctx, task)
}
