package task

import (
	"project-service/internal/service"
	pb "project-service/pkg/api/task_v1"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	taskService service.TaskService
}

func NewServer(taskService service.TaskService) *Server {
	return &Server{taskService: taskService}
}
