package controllers

import (
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
)

type CommandLineController struct {
	listTasksUseCaseInputPort         ports.ListTasksUseCaseInputPort
	createTaskUseCaseInputPort        ports.CreateTaskUseCaseInputPort
	completeTaskUseCaseInputPort      ports.CompleteTaskUseCaseInputPort
	listCompletedTaskUseCaseInputPort ports.ListCompletedTaskUseCaseInputPort
}

func NewCommandLineController(listTasksUseCaseInputPort ports.ListTasksUseCaseInputPort, createTaskUseCaseInputPort ports.CreateTaskUseCaseInputPort, completeTaskUseCaseInputPort ports.CompleteTaskUseCaseInputPort, listCompletedTaskUseCaseInputPort ports.ListCompletedTaskUseCaseInputPort) *CommandLineController {
	return &CommandLineController{listTasksUseCaseInputPort: listTasksUseCaseInputPort, createTaskUseCaseInputPort: createTaskUseCaseInputPort, completeTaskUseCaseInputPort: completeTaskUseCaseInputPort, listCompletedTaskUseCaseInputPort: listCompletedTaskUseCaseInputPort}
}

func (c *CommandLineController) ListCompleted() {
	c.listCompletedTaskUseCaseInputPort.Execute()
}

func (c *CommandLineController) ListTasks() {
	c.listTasksUseCaseInputPort.Execute()
}

func (c *CommandLineController) CreateTask(name, description string) {
	c.createTaskUseCaseInputPort.Execute(name, description)
}

func (c *CommandLineController) CompleteTask(id string) {
	c.completeTaskUseCaseInputPort.Execute(id)
}
