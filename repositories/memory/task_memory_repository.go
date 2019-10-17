package memory

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/qushot/go-cleanarchitecture-tasks/models"
	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
)

type taskMemoryRepository struct {
	tMap map[valueobject.TaskID]*models.Task
}

func NewTaskRepository() repositories.TaskRepository {
	return &taskMemoryRepository{
		tMap: map[valueobject.TaskID]*models.Task{},
	}
}

func (t *taskMemoryRepository) List() ([]*models.Task, error) {
	var tasks []*models.Task
	for _, task := range t.tMap {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *taskMemoryRepository) Get(id *valueobject.TaskID) (*models.Task, error) {
	if _, ok := t.tMap[*id]; !ok {
		return nil, errors.New(fmt.Sprintf("task not found. id=%s", id))
	}
	return t.tMap[*id], nil
}

func (t *taskMemoryRepository) Create(task *models.Task) (*models.Task, error) {
	id := uuid.New().String()
	taskID := valueobject.NewTaskID(id)
	registered := models.NewTask().TaskFactory.Of(taskID, task.Name(), task.Description())
	t.tMap[*taskID] = registered
	return registered, nil
}

func (t *taskMemoryRepository) Delete(id *valueobject.TaskID) error {
	delete(t.tMap, *id)
	return nil
}
