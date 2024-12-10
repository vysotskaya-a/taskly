package task

import (
	"project-service/internal/service"
	pb "project-service/pkg/api/task_v1"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	taskService service.TaskService
}

func NewTaskServer(taskService service.TaskService) *TaskServer {
	return &TaskServer{taskService: taskService}
}
