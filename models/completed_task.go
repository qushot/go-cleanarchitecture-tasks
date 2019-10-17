package models

type CompletedTask struct {
	Task
	CompletedTaskFactory completedTaskFactory
}

func newCompletedTask() *CompletedTask {
	return &CompletedTask{}
}

type completedTaskFactory struct {
}

func (*completedTaskFactory) CreateNewModel(task *Task) *CompletedTask {
	return &CompletedTask{
		Task: Task{
			id:          task.id,
			name:        task.name,
			description: task.description,
		},
	}
}
