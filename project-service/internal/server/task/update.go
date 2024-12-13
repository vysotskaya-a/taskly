package task

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/errorz"
	pb "project-service/pkg/api/task_v1"
)

func (s *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	err := s.taskService.Update(ctx, req)
	switch {
	case errors.Is(err, errorz.ErrUserIDNotSet):
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")

	case errors.Is(err, errorz.ErrTaskNotFound):
		return nil, status.Error(codes.NotFound, "Task with this ID not found.")

	case errors.Is(err, errorz.ErrProjectAccessForbidden):
		return nil, status.Error(codes.PermissionDenied, "Access denied for this task.")

	case err != nil:
		log.Error().Err(err).Msg("error while updating task")
		return nil, status.Error(codes.Internal, "Failed to update task.")
	}

	return &pb.UpdateTaskResponse{Success: true}, nil
}
