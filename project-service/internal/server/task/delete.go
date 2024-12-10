package task

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "project-service/pkg/api/task_v1"
)

func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.taskService.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to delete task: %v", err)
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}
