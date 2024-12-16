package project

import (
	"context"
	"google.golang.org/grpc/metadata"
	"project-service/internal/errorz"
)

func (s *Service) AddUser(ctx context.Context, newUserID, projectID string) error {
	var userID string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userID = md["user_id"][0]
	}
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
