package create

import (
	"GoRestSimpleApi/internal/task"
	Task "GoRestSimpleApi/internal/task"
	"context"
)

//New Task
type TaskCreateService struct {
	taskRepository Task.TaskSaveRepository
}

func NewTaskService(TaskRepository Task.TaskSaveRepository) TaskCreateService {
	return TaskCreateService{
		taskRepository: TaskRepository,
	}
}

func (s TaskCreateService) Save(ctx context.Context, task task.Task) error {

	return s.taskRepository.Save(ctx, task)
}
