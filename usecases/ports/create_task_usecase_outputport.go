package ports

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
)

type CreateTaskUseCaseOutputPort interface {
	ErrorOutputPort
	EmitTask(task *models.Task)
}
