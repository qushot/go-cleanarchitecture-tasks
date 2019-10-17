package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qushot/go-cleanarchitecture-tasks/models/valueobject"
)

type Task struct {
	id          valueobject.TaskID
	name        valueobject.TaskName
	description valueobject.TaskDescription
	TaskFactory taskFactory
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) ID() *valueobject.TaskID {
	return &t.id
}

func (t *Task) Name() *valueobject.TaskName {
	return &t.name
}

func (t *Task) Description() *valueobject.TaskDescription {
	return &t.description
}

func (t *Task) Validate() []error {
	var errs []error

	if t.Name() == nil {
		errs = append(errs, errors.New("タスク名は必須です。"))
	} else {
		if err := t.Name().Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if t.Description() != nil {
		if err := t.Description().Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) < 1 {
		return nil
	}

	return errs
}

func (t *Task) Complete() *CompletedTask {
	return newCompletedTask().CompletedTaskFactory.CreateNewModel(t)
}

func (t *Task) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("id=%s", t.ID()))
	sb.WriteString(fmt.Sprintf(", name=%s", t.Name()))
	sb.WriteString(fmt.Sprintf(", description=%s", t.Description()))

	return fmt.Sprintf("Task{%s}", sb.String())
}

type taskFactory struct {
}

func (*taskFactory) CreateModelForRegister(name *valueobject.TaskName, description *valueobject.TaskDescription) *Task {
	return &Task{
		name:        *name,
		description: *description,
	}
}

func (*taskFactory) Of(id *valueobject.TaskID, name *valueobject.TaskName, description *valueobject.TaskDescription) *Task {
	return &Task{
		id:          *id,
		name:        *name,
		description: *description,
	}
}
