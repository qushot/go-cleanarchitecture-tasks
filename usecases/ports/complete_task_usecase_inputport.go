package ports

type CompleteTaskUseCaseInputPort interface {
	Execute(id string)
}
