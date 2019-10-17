package impl

import (
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
)

type listCompletedTaskUseCase struct {
	outputPort ports.ListCompletedTaskUseCaseOutputPort
	repository repositories.CompletedTaskRepository
}

func NewListCompletedTaskUseCase(outputPort ports.ListCompletedTaskUseCaseOutputPort, repository repositories.CompletedTaskRepository) ports.ListCompletedTaskUseCaseInputPort {
	return &listCompletedTaskUseCase{outputPort: outputPort, repository: repository}
}

func (c *listCompletedTaskUseCase) Execute() {
	tasks, err := c.repository.List()
	if err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	c.outputPort.EmitCompletedTasks(tasks)
}
