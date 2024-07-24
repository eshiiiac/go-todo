package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// structure for a taks
type task struct {
	Title       string `json:"string"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// new task
/*func newTask(Title string, Description string, Status string) task {
	t := task{
		Title:       Title,
		Description: Description,
		Status:      Status,
	}
	return t
}*/

// to get user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), err
}

// to add ask
func addTask(id int) {
	for {
		reader := bufio.NewReader(os.Stdin)

		title, err := getInput("title: ", reader)
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

	}
}

func main() {
	id := 1
	addTask(id)
}
