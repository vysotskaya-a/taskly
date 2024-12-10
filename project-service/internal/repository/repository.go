package repository

import (
	"context"
	"project-service/internal/models"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) (string, error)
	GetByID(ctx context.Context, id string) (*models.Project, error)
	GetAllByUserID(ctx context.Context, userID string) ([]*models.Project, error)
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id string) error
}

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) (string, error)
	GetTasksByProjectID(ctx context.Context, projectID string) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id string) error
}
