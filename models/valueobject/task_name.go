package valueobject

import (
	"errors"
	"fmt"
)

type TaskName struct {
	value string
}

func NewTaskName(value string) *TaskName {
	return &TaskName{
		value: value,
	}
}

func (t *TaskName) Value() string {
	return t.value
}

func (t *TaskName) Validate() error {
	if t.value == "" {
		return errors.New("タスク名が空はダメです。")
	}
	if len(t.value) > 5 {
		return errors.New("タスク名は5文字以下である必要があります。")
	}
	return nil
}

func (t *TaskName) String() string {
	return fmt.Sprintf("TaskName{value='%s'}", t.value)
}
