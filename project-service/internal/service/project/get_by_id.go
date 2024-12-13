package project

import (
	"context"
	"project-service/internal/errorz"
	"project-service/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id string) (*models.Project, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, pUserID := range project.Users {
		if userID == pUserID {
			return project, nil
		}
	}

	return nil, errorz.ErrProjectAccessForbidden
}
