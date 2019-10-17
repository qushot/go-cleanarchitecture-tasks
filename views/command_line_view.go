package views

import (
	"fmt"

	"github.com/qushot/go-cleanarchitecture-tasks/presenters/viewmodels"
	"github.com/qushot/go-cleanarchitecture-tasks/presenters/views"
)

const (
	decorationReset = "\033[0m"
)

type commandLineView struct {
}

func NewCommandLineView() views.PrintView {
	return &commandLineView{}
}

func (c *commandLineView) Print(viewModel *viewmodels.PrintViewModel) {
	colorDecoration := c.toColorString(viewModel.GetMessageColor())
	fmt.Print(colorDecoration)
	for _, message := range viewModel.GetMessages() {
		fmt.Println(message)
	}
	fmt.Print(decorationReset)
}

func (c *commandLineView) toColorString(color viewmodels.Color) string {
	switch color {
	case viewmodels.Red:
		return "\033[0;31m"
	case viewmodels.Green:
		return "\033[0;32m"
	default:
		return ""
	}
}
