package ports

type CreateTaskUseCaseInputPort interface {
	Execute(name, description string)
}
