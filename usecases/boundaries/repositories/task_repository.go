package repositories

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
)

type TaskRepository interface {
	List() ([]*models.Task, error)
	Get(id *valueobject.TaskID) (*models.Task, error)
	Create(task *models.Task) (*models.Task, error)
	Delete(id *valueobject.TaskID) error
}
