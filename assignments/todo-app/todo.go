package main

import "fmt"

type Todo struct {
	Item      string
	completed bool
}

func formatCompleted(c bool) string {
	if c {
		return "yes"
	}
	return "no"
}

func (t Todo) Display() string {
	return fmt.Sprintf("Item: %q, Completed: %v", t.Item, formatCompleted(t.completed))
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

	todos := make([]Todo, 10)
	for i, item := range todoItems {
		todos[i] = Todo{Item: item}
	}

	for _, todo := range todos {
		fmt.Println(todo.Display())
	}
}
