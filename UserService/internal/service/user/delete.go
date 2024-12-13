package user

import (
	"context"
	"user-service/internal/errorz"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return errorz.ErrUserIDNotSet
	}

	if userID != id {
		return errorz.ErrUserAccessDenied
	}

	return s.userRepository.Delete(ctx, id)
}
