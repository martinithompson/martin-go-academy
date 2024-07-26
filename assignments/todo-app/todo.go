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

func (t *Todos) Add(item string) {
	todo := Todo{Item: item}
	t.Items = append(t.Items, todo)
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
	ts.Items = loadedTodos
}

func main() {
	fs := os.DirFS(".")
	newTodos := Todos{}
	newTodos.Load(fs, "todos.json")
	newTodos.PrintDescriptions(os.Stdout)
}
