package task

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/errorz"
	"project-service/internal/models"
	pb "project-service/pkg/api/task_v1"
)

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      models.TaskStatus(req.Status),
		ProjectId:   req.ProjectId,
		ExecutorId:  req.Executor,
		Deadline:    req.Deadline.AsTime(),
	}

	taskID, err := s.taskService.Create(ctx, task)
	switch {
	case errors.Is(err, errorz.ErrUserIDNotSet):
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")

	case errors.Is(err, errorz.ErrProjectNotFound):
		return nil, status.Error(codes.NotFound, "Project with this ID not found.")

	case errors.Is(err, errorz.ErrTaskAccessForbidden):
		return nil, status.Error(codes.PermissionDenied, "Access denied for this task.")

	case err != nil:
		return nil, status.Errorf(codes.Internal, "Failed to create task.")
	}

	return &pb.CreateTaskResponse{Id: taskID}, nil
}
