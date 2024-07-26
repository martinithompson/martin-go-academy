package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	items []Todo
}

func (t *Todos) Add(item string) {
	todo := Todo{Item: item}
	t.items = append(t.items, todo)
}

func (t Todos) PrintDescriptions(out io.Writer) {
	for _, todo := range t.items {
		fmt.Fprintln(out, todo.Description())
	}
}

func (ts Todos) Json() string {
	json, _ := json.Marshal(ts.items)
	return string(json)
}

func (ts Todos) Save(out io.Writer) error {
	_, err := out.Write([]byte(ts.Json()))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	todoItems := []string{
		"Wash the car",
		"Buy groceries",
		"Read book",
		"Clean the house",
		"Pay bills",
		"Call the bank",
		"Walk the dog",
		"Exercise",
		"Plan holiday",
		"Write report",
	}

	todos := Todos{}
	for _, item := range todoItems {
		todos.Add(item)
	}

	todos.PrintDescriptions(os.Stdout)
	outputFile, _ := os.Create("./output.json")
	todos.Save(outputFile)
}
