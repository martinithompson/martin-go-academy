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
	fmt.Fprintf(out, "\n\t*** To-Do Options ***\n\tPlease enter an option number to continue:\n\t\n\t1) Add a new to-do item\n\t2) View all to-dos\n\t3) Update a to-do item\n\t4) Delete a to-do item\n\t5) Exit\n")
}

func readOption(in *bufio.Reader, max int) (int, error) {
	fmt.Print("\n\t> ")
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
		addTodo(os.Stdin)
	case 2:
		viewTodos(os.Stdin)
	case 3:
		updateTodo(os.Stdin)
	case 4:
		fmt.Println("Delete a to-do")
		deleteTodo()
	default:
		fmt.Println("Goodbye")
	}
}

func addTodo(out io.Writer) {
	fmt.Fprintln(out, "\t*** Add a to-do ***")
	fmt.Fprint(out, "\tEnter the new to-do item name: ")
	item, _ := readItem(reader)
	todoItems.AddTaskItems(item)
}

func viewTodos(out io.Writer) {
	if len(todoItems.Items) > 0 {
		fmt.Fprintf(out, "\t*** Your to-do list ***\n\n")
		todoItems.PrintDescriptions(os.Stdout)
	} else {
		fmt.Println("\tYour to-do list is empty :-)")
	}
}

func updateTodo(out io.Writer) {
	if len(todoItems.Items) > 0 {
		fmt.Fprintln(out, "\t*** Update a to-do ***")
		fmt.Fprintf(out, "\tEnter the number of the item to update:\n")
		todoItems.PrintDescriptions(os.Stdout)

		item, _ := readOption(reader, len(todoItems.Items))

		fmt.Fprintln(out, "\t1) Update item name")
		fmt.Fprintln(out, "\t2) Toggle completed status")

		updateOption, _ := readOption(reader, 2)

		if updateOption == 1 {
			fmt.Fprint(out, "\tEnter the updated to-do item name: ")
			updatedItem, _ := readItem(reader)
			todoItems.Items[item-1].Item = updatedItem
		} else {
			todoItems.Items[item-1].Completed = !todoItems.Items[item-1].Completed
		}

	} else {
		fmt.Println("Your todo list is empty :-)")
	}
}

func deleteTodo() {
	fmt.Println("Enter the number of the item to delete:")
	todoItems.PrintDescriptions(os.Stdout)

	item, _ := readOption(reader, len(todoItems.Items))
	todoItems.Items = append(todoItems.Items[:item-1], todoItems.Items[item:]...)
}

func main() {

	var option int
	var err error
	for option != 5 {
		menu(os.Stdout)
		for {
			option, err = readOption(reader, 5)
			if err != nil {
				fmt.Println("Invalid option, please enter 1-5:")
			} else {
				break
			}
		}

		fmt.Printf("\tYou selected option: %d\n\n", option)
		handleOption(option)
	}
}
