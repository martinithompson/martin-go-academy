package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	todos "todo-app/project/todos"
)

var todoItems = todos.Todos{}
var reader = bufio.NewReader(os.Stdin)

func menu(out io.Writer) {
	fmt.Fprintf(out, "\t*** To-Do Options***\n\tPlease enter an option number to continue:\n\t\n\t1: Add a new to-do item\n\t2: View all todos\n\t3: Update a todo item\n\t4: Delete a todo item\n\t5: Exit\n")
}

func readOption(in *bufio.Reader, max int) (int, error) {
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	i, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	if i < 1 || i > max {
		return 0, errors.New("invalid option")
	}
	return i, nil
}

func readItem(in *bufio.Reader) (string, error) {
	item, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	item = strings.TrimSpace(item)
	return item, nil
}

func handleOption(option int) {
	switch option {
	case 1:
		fmt.Println("Add a todo")
		addTodo(os.Stdin)
	case 2:
		fmt.Println("View todos")
		viewTodos()
	case 3:
		fmt.Println("Update a todo")
		updateTodo()
	case 4:
		fmt.Println("Delete a todo")
		deleteTodo()
	default:
		fmt.Println("Goodbye")
	}
}

func addTodo(out io.Writer) {
	fmt.Fprintln(out, "Enter the task name: ")
	item, _ := readItem(reader)
	todoItems.AddTaskItems(item)
}

func viewTodos() {
	todoItems.PrintDescriptions(os.Stdout)
}

func updateTodo() {
	fmt.Println("To do")
}

func deleteTodo() {
	fmt.Println("To do")
}

func main() {

	menu(os.Stdout)
	var option int
	var err error
	for option != 5 {
		for {
			option, err = readOption(reader, 5)
			if err != nil {
				fmt.Println("Invalid option, please enter 1-5:")
			} else {
				break
			}
		}

		fmt.Printf("You selected option: %d\n", option)
		handleOption(option)
	}
}
