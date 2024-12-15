package project

import (
	"context"
	"project-service/internal/errorz"
)

func (s *Service) SubscribeOnNotifications(ctx context.Context, projectID string, telegramID int64) error {
	userID, ok := ctx.Value("user_id").(string)
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
