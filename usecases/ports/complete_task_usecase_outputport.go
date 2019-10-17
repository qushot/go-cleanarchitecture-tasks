package ports

type CompleteTaskUseCaseOutputPort interface {
	ErrorOutputPort
	EmitCompletedTaskID(completedTaskID string)
}
