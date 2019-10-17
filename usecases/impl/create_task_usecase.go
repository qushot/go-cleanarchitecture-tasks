package impl

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
)

type createTaskUseCase struct {
	outputPort ports.CreateTaskUseCaseOutputPort
	repository repositories.TaskRepository
}

func NewCreateTaskUseCase(outputPort ports.CreateTaskUseCaseOutputPort, repository repositories.TaskRepository) ports.CreateTaskUseCaseInputPort {
	return &createTaskUseCase{outputPort: outputPort, repository: repository}
}

func (c *createTaskUseCase) Execute(name, description string) {
	taskName := valueobject.NewTaskName(name)
	taskDescription := valueobject.NewTaskDescription(description)

	task := models.NewTask().TaskFactory.CreateModelForRegister(taskName, taskDescription)
	if errs := task.Validate(); errs != nil {
		c.outputPort.ShowError(errs)
		return
	}

	registeredTask, err := c.repository.Create(task)
	if err != nil {
		c.outputPort.ShowError([]error{err})
		return
	}

	c.outputPort.EmitTask(registeredTask)
}
