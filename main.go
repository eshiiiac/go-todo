package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// structure for a taks
type task struct {
	Title       string `json:"string"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// new task
func newTask(Title string, Description string, Status string) task {
	t := task{
		Title:       Title,
		Description: Description,
		Status:      Status,
	}
	return t
}

// to get user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), err
}

func getIntInput(prompt string, r *bufio.Reader) (int, error) {
    fmt.Print(prompt)
    input, err := r.ReadString('\n')
    if err != nil {
        return 0, err
    }
    input = strings.TrimSpace(input)
    value, err := strconv.Atoi(input)
    if err != nil {
        return 0, err
    }
    return value, nil
}


func addTask(id int) {
	for {
		reader := bufio.NewReader(os.Stdin)

		title, err := getInput("title: ", reader)
		if err != nil {
			fmt.Println("error in adding todo")
		}
		description, err := getInput("description: ", reader)
		if err != nil {
			fmt.Println("error in adding todo")
		}
		//nested map to store multiple variables in a single map(map inside a map)
		task := make(map[int]map[string]string)
		//id being the main map, map2(where title,desc and status is stored)
		task[id] = map[string]string{
			"title":       title,
			"description": description,
			"status":      "pending",
		}
		fmt.Printf("task %v added, ctrl+c to save\n", id)
		id++ // task id incremented
		// storing the task in todo file
		jsonData, err := json.Marshal(task)
		if err != nil {
			fmt.Println("Error sterlizing map ", err)
			return
		}

		//creating a new file, if exists the data is appended
		file, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("error creating file ", err)
			return
		}
		defer file.Close()
		//  append additional data to the file if needed
		_, err = file.Write([]byte("\n")) //adding a new line
		if err != nil {
			fmt.Println("Error writing newline to file:", err)
			return
		}

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("Error appending data:", err)
			return
		}

		exit, err := getInput("add next task(Y or N): ", reader)
		if err != nil {
			fmt.Println("Error appending data:", err)
			return
		}
		switch exit {
		case "y":
			continue
		default:
			main()
		}
	}
}

func readTask() {
	fmt.Println("Your Tasks:")
	content, err := os.ReadFile("todos.txt")
	if err != nil {
		fmt.Println("ERR ", err)
		return
	}
	fmt.Println(string(content))
	main()
}

func updateTask(status string) {

	reader := bufio.NewReader(os.Stdin)
	taskNumStr, err := getInput("enter task no: ", reader)
	if err != nil {
		fmt.Println("ERR ", err)
		return
	}

	taskNum, err := strconv.Atoi(taskNumStr)
	if err != nil {
		fmt.Println("Invalid task number:", err)
		return
	}

	task := make(map[int]string)
	updateStatus, err := getInput("task completed?(y | n) ", reader)
	if err != nil {
		fmt.Println("ERR ", err)
		return
	}

	switch updateStatus {
	case "y":
		status := "completed"
		task[taskNum] = status
		main()
		return
	default:
		status := "pending"
		fmt.Println(status)
		main()
		return
	}

}

func deleteTask() {
	var id int
	reader := bufio.NewReader(os.Stdin)
	id, err := getIntInput("Enter task no. to delete: ", reader)
	if err != nil {
		fmt.Println("err deleting task")
		return
	}
	task := make(map[int]string)

	if _, exists := task[id]; exists {
		delete(task, id)
		fmt.Printf("Task %d has been deleted.\n", id)
    } else {
        fmt.Printf("Task %d not found.\n", id)
}
}
func main() {
	id := 1
	var status string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n1. Add Task\n2. Show Tasks\n3. Update Task\n4. Delete Task")
	userChoice, err := getInput("your choice: ", reader)
	if err != nil {
		fmt.Println("invalid user choice ", err)
		return
	}

	switch userChoice {
	case "1":
		addTask(id)
	case "2":
		readTask()
	case "3":
		updateTask(status)
	case "4":
		deleteTask()
	default:
		fmt.Println("invalid user choice ", err)
		main()
	}
}
