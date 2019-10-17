package valueobject

import (
	"errors"
	"fmt"
)

type TaskDescription struct {
	value string
}

func NewTaskDescription(value string) *TaskDescription {
	return &TaskDescription{
		value: value,
	}
}

func (t *TaskDescription) Value() string {
	return t.value
}

func (t *TaskDescription) Validate() error {
	if t.value == "" {
		return errors.New("タスクの説明が空はダメです。")
	}
	if len(t.value) > 10 {
		return errors.New("タスクの説明は10文字以下である必要があります。")
	}
	return nil
}

func (t *TaskDescription) String() string {
	return fmt.Sprintf("TaskDescription{value='%s'}", t.value)
}
