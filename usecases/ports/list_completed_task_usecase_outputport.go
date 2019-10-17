package ports

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
)

type ListCompletedTaskUseCaseOutputPort interface {
	ErrorOutputPort
	EmitCompletedTasks(tasks []*models.CompletedTask)
}
