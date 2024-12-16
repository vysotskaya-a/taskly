package task

import (
	"context"
	"google.golang.org/grpc/metadata"
	"project-service/internal/errorz"
	"project-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.Task, error) {
	var userID string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userID = md["user_id"][0]
	}
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
