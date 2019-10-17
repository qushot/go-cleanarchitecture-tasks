package views

import (
	"github.com/qushot/go-cleanarchitecture-tasks/presenters/viewmodels"
)

type PrintView interface {
	Print(viewModel *viewmodels.PrintViewModel)
}
