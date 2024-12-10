package task

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/models"
	pb "project-service/pkg/api/task_v1"
)

func (s *TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	id, err := s.taskService.CreateTask(ctx, &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		ProjectId:   req.ProjectId,
		ExecutorId:  req.Executor,
		Deadline:    req.Deadline.AsTime(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
	}

	return &pb.CreateTaskResponse{Id: id}, nil
}
