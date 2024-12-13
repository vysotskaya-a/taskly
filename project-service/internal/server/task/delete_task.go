package task

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/errorz"
	pb "project-service/pkg/api/task_v1"
)

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.taskService.Delete(ctx, req.Id)
	switch {
	case errors.Is(err, errorz.ErrUserIDNotSet):
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")

	case errors.Is(err, errorz.ErrTaskNotFound):
		return nil, status.Error(codes.NotFound, "Task with this ID not found.")

	case errors.Is(err, errorz.ErrTaskAccessForbidden):
		return nil, status.Error(codes.PermissionDenied, "Access denied for this task.")

	case err != nil:
		return nil, status.Errorf(codes.Internal, "Failed to delete task.")
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}
