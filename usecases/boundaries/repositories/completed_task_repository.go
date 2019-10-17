package repositories

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
)

type CompletedTaskRepository interface {
	List() ([]*models.CompletedTask, error)
	Create(task *models.CompletedTask) (*models.CompletedTask, error)
}
