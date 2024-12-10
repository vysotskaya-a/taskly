package task

import "project-service/internal/repository"

type ServiceTask struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) *ServiceTask {
	return &ServiceTask{taskRepo: taskRepo}
}
