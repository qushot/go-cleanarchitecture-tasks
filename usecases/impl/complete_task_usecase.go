package impl

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
)

type completeTaskUseCase struct {
	outputPort              ports.CompleteTaskUseCaseOutputPort
	taskRepository          repositories.TaskRepository
	completedTaskRepository repositories.CompletedTaskRepository
}

func NewCompleteTaskUseCase(outputPort ports.CompleteTaskUseCaseOutputPort, taskRepository repositories.TaskRepository, completedTaskRepository repositories.CompletedTaskRepository) ports.CompleteTaskUseCaseInputPort {
	return &completeTaskUseCase{outputPort: outputPort, taskRepository: taskRepository, completedTaskRepository: completedTaskRepository}
}

func (c *completeTaskUseCase) Execute(id string) {
	taskID := valueobject.NewTaskID(id)
	task, err := c.taskRepository.Get(taskID)
	if err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	completedTask := task.Complete()
	if err := c.taskRepository.Delete(taskID); err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	if _, err := c.completedTaskRepository.Create(completedTask); err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	c.outputPort.EmitCompletedTaskID(completedTask.ID().Value())
}
