package presenters

import (
	"fmt"

	"github.com/qushot/go-cleanarchitecture-tasks/models"
	"github.com/qushot/go-cleanarchitecture-tasks/presenters/viewmodels"
	"github.com/qushot/go-cleanarchitecture-tasks/presenters/views"
)

type PrintPresenter struct {
	view views.PrintView
}

func NewPrintPresenter(view views.PrintView) *PrintPresenter {
	return &PrintPresenter{view: view}
}

func (p *PrintPresenter) EmitTask(task *models.Task) {
	messages := []string{
		"Following task was created!!",
		task.String(),
	}
	viewModel := viewmodels.NewPrintViewModel(messages, viewmodels.Green)
	p.view.Print(viewModel)
}

func (p *PrintPresenter) EmitCompletedTaskID(completedTaskID string) {
	messages := []string{
		fmt.Sprintf("WELL DONE. task id: %s changed to complete. :D", completedTaskID),
	}
	viewModel := viewmodels.NewPrintViewModel(messages, viewmodels.Green)
	p.view.Print(viewModel)
}

func (p *PrintPresenter) EmitTasks(tasks []*models.Task) {
	var messages []string
	if len(tasks) < 1 {
		messages = append(messages, "no tasks")
	} else {
		for _, task := range tasks {
			messages = append(messages, task.String())
		}
	}

	viewModel := viewmodels.NewPrintViewModel(messages, viewmodels.Green)
	p.view.Print(viewModel)
}

func (p *PrintPresenter) EmitCompletedTasks(tasks []*models.CompletedTask) {
	var messages []string
	if len(tasks) < 1 {
		messages = append(messages, "no tasks")
	} else {
		for _, task := range tasks {
			messages = append(messages, task.String())
		}
	}

	viewModel := viewmodels.NewPrintViewModel(messages, viewmodels.Green)
	p.view.Print(viewModel)
}
func (p *PrintPresenter) ShowError(errs []error) {
	var messages []string
	if 0 < len(errs) {
		for _, err := range errs {
			messages = append(messages, fmt.Sprintf("Error: %s", err.Error()))
		}
	}
	viewModel := viewmodels.NewPrintViewModel(messages, viewmodels.Red)
	p.view.Print(viewModel)
}
