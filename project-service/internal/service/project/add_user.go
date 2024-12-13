package project

import (
	"context"
	"project-service/internal/errorz"
)

func (s *Service) AddUser(ctx context.Context, newUserID, projectID string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, projectID)
	if err != nil {
		return err
	}

	if userID != project.AdminID {
		return errorz.ErrProjectAccessForbidden
	}

	project.Users = append(project.Users, newUserID)

	return s.projectRepository.Update(ctx, project)
}
