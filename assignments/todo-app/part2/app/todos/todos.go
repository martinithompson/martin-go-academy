package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
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

func (ts *Todos) AddTaskItems(items ...string) {
	for _, item := range items {
		ts.Items = append(ts.Items, Todo{Item: item})
	}
}

func (ts *Todos) AddTodoItems(items ...Todo) {
	ts.Items = append(ts.Items, items...)
}

func (ts *Todos) DeleteTodoItem(index int) error {
	if len(ts.Items) < index {
		return errors.New("index out of range")
	}
	ts.Items = append(ts.Items[:index-1], ts.Items[index:]...)
	return nil
}

func (t Todos) PrintDescriptions(out io.Writer) {
	for i, todo := range t.Items {
		fmt.Fprintf(out, "\t%d> %s\n", i+1, todo.Description())
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
