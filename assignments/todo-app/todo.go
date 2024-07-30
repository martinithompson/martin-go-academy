package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

type Todo struct {
	Item      string
	Completed bool
}

func formatCompleted(c bool) string {
	if c {
		return "yes"
	}
	return "no"
}

func (t Todo) Description() string {
	return fmt.Sprintf("Item: %q, Completed: %v", t.Item, formatCompleted(t.Completed))
}

type Todos struct {
	Items []Todo
}

func (t *Todos) AddTaskItems(items ...string) {
	for _, item := range items {
		t.Items = append(t.Items, Todo{Item: item})
	}
}

func (t *Todos) AddTodoItems(items ...Todo) {
	t.Items = append(t.Items, items...)
}

func (t Todos) PrintDescriptions(out io.Writer) {
	for _, todo := range t.Items {
		fmt.Fprintln(out, todo.Description())
	}
}

func (ts Todos) Json() string {
	json, _ := json.Marshal(ts.Items)
	return string(json)
}

func (ts Todos) Save(out io.Writer) error {
	_, err := out.Write([]byte(ts.Json()))
	if err != nil {
		return err
	}
	return nil
}

func (ts *Todos) Load(fileSystem fs.FS, fileName string) {
	todosFile, err := fileSystem.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer todosFile.Close()

	content, err := io.ReadAll(todosFile)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var loadedTodos []Todo
	jsonErr := json.Unmarshal([]byte(content), &loadedTodos)
	if jsonErr != nil {
		fmt.Println("Error unmarshalling JSON:", jsonErr)
	}
	fmt.Println(loadedTodos)
	ts.AddTodoItems(loadedTodos...)
}

var tasksToAdd = []string{"go shopping", "wash the car", "walk the dog", "do laundry", "pay bills",
	"clean the house", "cook dinner", "read a book", "exercise", "call bank"}

func main() {
	// TODO - menu switch statement to choose exercise
	// exerciseOne()
	// exerciseTwo()
	exerciseThree()
	// exerciseFour()
}

// Print a list of 10 things to do
func exerciseOne() {
	todos := Todos{}
	todos.AddTaskItems(tasksToAdd...)
	todos.PrintDescriptions(os.Stdout)
}

// Display list of 10 things in json
func exerciseTwo() {
	todos := Todos{}
	todos.AddTaskItems(tasksToAdd...)
	fmt.Println(todos.Json())
}

func exerciseThree() {
	todos := Todos{}
	todos.AddTaskItems(tasksToAdd...)
	outputFile, _ := os.Create("./output.json")
	todos.Save(outputFile)
}

// read a todo list from a json file
func exerciseFour() {
	fs := os.DirFS(".")
	newTodos := Todos{}
	newTodos.Load(fs, "todos.json")
	newTodos.PrintDescriptions(os.Stdout)
}
