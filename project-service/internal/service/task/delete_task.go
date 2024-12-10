package task

import "context"

func (s *ServiceTask) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepo.DeleteTask(ctx, id)
}
