package task

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/models"
	pb "project-service/pkg/api/task_v1"
)

func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task := &models.Task{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		ExecutorId:  req.Executor,
		Deadline:    req.Deadline.AsTime(),
	}

	err := s.taskService.UpdateTask(ctx, task)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to update task: %v", err)
	}

	return &pb.UpdateTaskResponse{Success: true}, nil
}
