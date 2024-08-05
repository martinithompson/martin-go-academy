package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"

	"github.com/google/uuid"
)

type Todo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (t Todo) Description() string {
	return fmt.Sprintf("Name: %q, Completed: %v", t.Name, formatCompleted(t.Completed))
}

type TodoList struct {
	Items []Todo
}

func (ts *TodoList) AddTodosByName(names ...string) {
	// if adding new todo by name only, generate a new id and leave completed as false
	for _, name := range names {
		ts.Items = append(ts.Items, Todo{Name: name, Id: uuid.New().String()})
	}
}

func (ts *TodoList) AddTodos(items ...Todo) {
	ts.Items = append(ts.Items, items...)
}

func findTodo(todos []Todo, id string, action func(i int) error) error {
	for i, todo := range todos {
		if todo.Id == id {
			return action(i)
		}
	}
	return errors.New("todo not found")
}

func (ts *TodoList) DeleteTodo(id string) error {
	return findTodo(ts.Items, id, func(i int) error {
		ts.Items = append(ts.Items[:i], ts.Items[i+1:]...)
		return nil
	})
}

func (ts *TodoList) UpdateTodo(updated Todo) error {
	return findTodo(ts.Items, updated.Id, func(i int) error {
		ts.Items[i].Name = updated.Name
		ts.Items[i].Completed = updated.Completed
		return nil
	})
}

func (ts TodoList) PrintDescriptions(out io.Writer) {
	for i, todo := range ts.Items {
		fmt.Fprintf(out, "\t%d> %s\n", i+1, todo.Description())
	}
}

func (ts TodoList) Json() string {
	json, _ := json.Marshal(ts.Items)
	return string(json)
}

func (ts TodoList) Save(out io.Writer) error {
	_, err := out.Write([]byte(ts.Json()))
	if err != nil {
		return err
	}
	return nil
}

func (ts *TodoList) Load(fileSystem fs.FS, fileName string) {
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
	ts.AddTodos(loadedTodos...)
}

func formatCompleted(c bool) string {
	if c {
		return "yes"
	}
	return "no"
}
