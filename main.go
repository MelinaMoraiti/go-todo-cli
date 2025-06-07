package main

import (
	"bufio"
    "encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string
	Done        bool
}

var tasks []Task
const taskFile = "tasks.json"

func main() {
	loadTasks()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Go To-Do CLI!")
	for {
		fmt.Print("\nCommands: add <task>, done <task number>, list, quit\n> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := strings.SplitN(input, " ", 2)
		cmd := args[0]

		switch cmd {
		case "add":
			if len(args) < 2 {
				fmt.Println("Please provide a task description.")
				continue
			}
			addTask(args[1])
		case "list":
			listTasks()
		case "done":
			if len(args) < 2 {
				fmt.Println("Please provide a task description.")
				continue
			}
			num, err := strconv.Atoi(args[1])
				if err != nil || num < 1 || num > len(tasks) {
					fmt.Println("Invalid task number.")
					continue
				}
			completeTask(num - 1) 
		case "quit":
			saveTasks()
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}

func addTask(desc string) {
	task := Task{Description: desc, Done: false}
	tasks = append(tasks, task)
	fmt.Println("Added:", desc)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

func completeTask(i int) {
	tasks[i].Done = true
	saveTasks()
	fmt.Println("Task marked as done:", tasks[i].Description)
}

func saveTasks() {
	file, err := os.Create(taskFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}

func loadTasks() {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return
		}
		fmt.Println("Error opening tasks file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
	}
}