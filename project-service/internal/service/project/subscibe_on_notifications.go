package project

import (
	"context"
	"google.golang.org/grpc/metadata"
	"project-service/internal/errorz"
)

func (s *Service) SubscribeOnNotifications(ctx context.Context, projectID string, telegramID int64) error {
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

	for _, pUserID := range project.Users {
		if userID == pUserID {
			project.NotificationSubscribersTGIDS = append(project.NotificationSubscribersTGIDS, telegramID)
			return s.projectRepository.Update(ctx, project)
		}
	}

	return errorz.ErrProjectAccessForbidden
}
