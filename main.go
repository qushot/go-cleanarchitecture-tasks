package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/qushot/go-cleanarchitecture-tasks/controllers"
	"github.com/qushot/go-cleanarchitecture-tasks/presenters"
	"github.com/qushot/go-cleanarchitecture-tasks/presenters/views"
	"github.com/qushot/go-cleanarchitecture-tasks/repositories/memory"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/boundaries/repositories"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/impl"
	"github.com/qushot/go-cleanarchitecture-tasks/usecases/ports"
	clviews "github.com/qushot/go-cleanarchitecture-tasks/views"
)

var (
	taskRepository          repositories.TaskRepository          = memory.NewTaskRepository()
	completedTaskRepository repositories.CompletedTaskRepository = memory.NewCompletedTaskRepository()
	printView               views.PrintView                      = clviews.NewCommandLineView()
	printPresenter          *presenters.PrintPresenter           = presenters.NewPrintPresenter(printView)

	listTasksUseCaseInputPort         ports.ListTasksUseCaseInputPort         = impl.NewListTaskUseCase(taskRepository, printPresenter)
	createTaskUseCaseInputPort        ports.CreateTaskUseCaseInputPort        = impl.NewCreateTaskUseCase(printPresenter, taskRepository)
	completeTaskUseCaseInputPort      ports.CompleteTaskUseCaseInputPort      = impl.NewCompleteTaskUseCase(printPresenter, taskRepository, completedTaskRepository)
	listCompletedTaskUseCaseInputPort ports.ListCompletedTaskUseCaseInputPort = impl.NewListCompletedTaskUseCase(printPresenter, completedTaskRepository)

	controller = controllers.NewCommandLineController(listTasksUseCaseInputPort, createTaskUseCaseInputPort, completeTaskUseCaseInputPort, listCompletedTaskUseCaseInputPort)
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Print("input > ")
		if sc.Scan() {
			input := sc.Text()
			if input == "exit" {
				fmt.Println("Bye!")
				break
			}

			inputs := strings.Split(input, " ")
			if len(inputs) < 1 {
				fmt.Println("1st argument is required.")
				continue
			}

			resource := inputs[0]
			switch resource {
			case "help":
				fmt.Println("sorry, this functions is not implemented.")
				continue
			case "tasks":
				if len(inputs) < 2 {
					fmt.Println("action is required")
					break
				}
				action := inputs[1]
				switch action {
				case "list":
					if len(inputs) > 2 {
						flag := inputs[2]
						switch flag {
						case "-c":
							controller.ListCompleted()
							continue
						}
					}
					controller.ListTasks()
					break
				case "create":
					if len(inputs) < 4 {
						fmt.Println("id is required as 3rd argument.")
						break
					}
					controller.CreateTask(inputs[2], inputs[3])
					break
				case "complete":
					if len(inputs) < 3 {
						fmt.Println("id is required as 3rd argument.")
						break
					}
					controller.CompleteTask(inputs[2])
					break
				default:
					fmt.Println("action is required.")
					break
				}
				break
			default:
				fmt.Println("nothing to do. :D")
				break
			}
		}

	}
}
