package read

import (
	Task "GoRestSimpleApiLimpio/internal/task"
)

//Task Get ID
type TaskGetIdService struct {
	taskRepository Task.TaskGetIdRepository
}

func NewTaskServiceID(TaskRepository Task.TaskGetIdRepository) TaskGetIdService {
	return TaskGetIdService{
		taskRepository: TaskRepository,
	}
}

func (s TaskGetIdService) ReadID(Id string) ([]Task.TaskDB, error) {

	return s.taskRepository.ReadID(Id)
}
