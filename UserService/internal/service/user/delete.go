package user

import "context"

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.userRepository.Delete(ctx, id)
}
