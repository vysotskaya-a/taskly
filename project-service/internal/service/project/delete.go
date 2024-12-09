package project

import "context"

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.projectRepository.Delete(ctx, id)
}
