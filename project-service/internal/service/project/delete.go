package project

import (
	"context"
	"project-service/internal/errorz"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if userID != project.AdminID {
		return errorz.ErrProjectAccessForbidden
	}

	return s.projectRepository.Delete(ctx, id)
}
