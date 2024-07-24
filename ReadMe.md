### explaining the future me about the garbage code i wrote _(few were copied from chatgpt(reference))_

beginning with the 
- created a structure for *task*
```
type task struct {
	Title       string `json:"string"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
```

## created a function that takes user input

```
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), err
}
```
- bufio is a library
- the ``r.ReadString('\n')`` makes the pprogram to store the data after user hits enter/ new line
- the ``strings.TrimSpace(input)`` trims the white spaces in user input
> i could've used the default *scan* to store data but calling function was a beter approach

## adding a new task

```
func addTask(id int) {
	for {
		reader := bufio.NewReader(os.Stdin)

		title, err := getInput("title: ", reader)
		description, err := getInput("description: ", reader)
		if err != nil {
			fmt.Println("error in adding todo")
		}
		
		task := make(map[int]map[string]string)
		//id being the main map, map2(where title,desc and status is stored)
		task[id] = map[string]string{
			"title":       title,
			"description": description,
			"status":      "pending",
		}
		fmt.Printf("task %v added, ctrl+c to save\n", id)
		id++ 
```
- ``for{}`` an infinite for loop
- ``reader := bufio.NewReader(os.Stdin)`` for storing user input 

example:
>`` title, err := getInput("title: ", reader)
		description, err := getInput("description: ", reader)``
data is stored to the variable *title*

## storing the data into a file
```
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
```

## main function
```
func main() {
	id := 1
	addTask(id)
}
```