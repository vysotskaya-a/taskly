package project

import (
	"context"
	"google.golang.org/grpc/metadata"
	"project-service/internal/errorz"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	var userID string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userID = md["user_id"][0]
	}
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
