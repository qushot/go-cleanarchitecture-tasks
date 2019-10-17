package ports

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
)

type ListTasksUseCaseOutputPort interface {
	ErrorOutputPort
	EmitTasks(tasks []*models.Task)
}
