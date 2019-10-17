package ports

type ErrorOutputPort interface {
	ShowError(errs []error)
}
