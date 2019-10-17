package memory

import (
	"github.com/qushot/go-cleanarchitecture-tasks/models"
	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
)

type completedTaskMemoryRepository struct {
	ctMap map[valueobject.TaskID]*models.CompletedTask
}

func NewCompletedTaskRepository() repositories.CompletedTaskRepository {
	return &completedTaskMemoryRepository{
		ctMap: map[valueobject.TaskID]*models.CompletedTask{},
	}
}

func (t *completedTaskMemoryRepository) List() ([]*models.CompletedTask, error) {
	var tasks []*models.CompletedTask
	for _, task := range t.ctMap {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *completedTaskMemoryRepository) Create(task *models.CompletedTask) (*models.CompletedTask, error) {
	taskID := task.ID()
	t.ctMap[*taskID] = task
	return task, nil
}
