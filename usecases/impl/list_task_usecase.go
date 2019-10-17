package impl

import (
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
)

type listTaskUseCase struct {
	repository repositories.TaskRepository
	outputPort ports.ListTasksUseCaseOutputPort
}

func NewListTaskUseCase(repository repositories.TaskRepository, outputPort ports.ListTasksUseCaseOutputPort) ports.ListTasksUseCaseInputPort {
	return &listTaskUseCase{repository: repository, outputPort: outputPort}
}

func (c *listTaskUseCase) Execute() {
	tasks, err := c.repository.List()
	if err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	c.outputPort.EmitTasks(tasks)
}
