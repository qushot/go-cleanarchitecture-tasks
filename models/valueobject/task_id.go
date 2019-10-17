package valueobject

import "fmt"

type TaskID struct {
	value string
}

func NewTaskID(value string) *TaskID {
	return &TaskID{
		value: value,
	}
}

func (t *TaskID) Value() string {
	return t.value
}

func (t *TaskID) String() string {
	return fmt.Sprintf("TaskID{value='%s'}", t.value)
}
