package tasks

import "fmt"

const (
	AddCommand      = "add"
	RemoveCommand   = "remove"
	CompleteCommand = "complete"
	ListCommand     = "list"
	Clear           = "clear"
	ExitCommand     = "exit"
)

func Run() {
	tasks := []string{}

	for {
		fmt.Print("Enter command (add [TEXT], remove [Number], complete [Number], list, clear : cleans the screen , or exit): ")
		userText := readText()
		if userText == "" {
			fmt.Println("Please enter a command.")
			continue
		}

		command := getFirstWord(userText)

		switch command {
		case AddCommand:
			task := removeFirstWord(userText)
			if task == "" {
				fmt.Println("Please specify a task to add.")
				continue
			}
			tasks = addTask(tasks, userText)

		case ListCommand:
			listTasks(tasks)

		case RemoveCommand:
			taskNumber, err := parseTaskNumber(userText)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task number.")
				continue
			}
			var removeErr error
			tasks, removeErr = removeTask(tasks, taskNumber)
			if removeErr != nil {
				fmt.Println(removeErr)
				continue
			}
			fmt.Println("Task removed successfully!")

		case CompleteCommand:
			taskNumber, err := parseTaskNumber(userText)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task number.")
				continue
			}
			var completeErr error
			tasks, completeErr = completeTask(tasks, taskNumber)
			if completeErr != nil {
				fmt.Println(completeErr)
				continue
			}
			fmt.Printf("Task '%s' marked complete.\n", tasks[taskNumber-1])

		case ExitCommand:
			fmt.Println("Exiting...")
			return
		case Clear:
			clearScreen()

		default:
			fmt.Println("Invalid command. Please enter add, remove, complete, list, or exit.")
		}
	}
}
