package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func readText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(scanner.Text())
}

func removeFirstWord(input string) string {
	words := strings.Fields(input)
	if len(words) > 1 {
		return strings.Join(words[1:], " ")
	}
	return ""
}

func getFirstWord(input string) string {
	words := strings.Fields(input)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

func addTask(tasks []string, input string) []string {
	return append(tasks, removeFirstWord(input))
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func removeTask(tasks []string, taskNumber int) ([]string, error) {
	if taskNumber < 1 || taskNumber > len(tasks) {
		return tasks, fmt.Errorf("task number out of range")
	}
	return append(tasks[:taskNumber-1], tasks[taskNumber:]...), nil
}

func completeTask(tasks []string, taskNumber int) ([]string, error) {
	if taskNumber < 1 || taskNumber > len(tasks) {
		return tasks, fmt.Errorf("task number out of range")
	}
	tasks[taskNumber-1] = "**COMPLETED** - " + tasks[taskNumber-1]
	return tasks, nil
}

func parseTaskNumber(input string) (int, error) {
	return strconv.Atoi(removeFirstWord(input))
}

func clearScreen() {
	var clearCommand *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		clearCommand = exec.Command("clear")
	case "windows":
		clearCommand = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Unsupported platform!")
		return
	}
	clearCommand.Stdout = os.Stdout
	clearCommand.Run()
}
