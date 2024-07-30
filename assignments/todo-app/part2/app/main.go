package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// var tasksToAdd = []string{"go shopping", "wash the car", "walk the dog", "do laundry", "pay bills",
// 	"clean the house", "cook dinner", "read a book", "exercise", "call bank"}

func menu(out io.Writer) {
	fmt.Fprintf(out, "\t*** To-Do Options***\n\tPlease enter an option number to continue:\n\t\n\t1: Add a new to-do item\n\t2: View all todos\n\t3: Update a todo item\n\t4: Delete a todo item\n")
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	menu(os.Stdout)
	var option int
	var err error
	for {
		option, err = readOption(reader, 4)
		if err != nil {
			fmt.Println("Invalid option, please enter 1-4:")
		} else {
			break
		}
	}

	fmt.Printf("You selected option: %d\n", option)
}

// Print a list of 10 things to do
// func exerciseOne() {
// 	todos := todos.Todos{}
// 	todos.AddTaskItems(tasksToAdd...)
// 	todos.PrintDescriptions(os.Stdout)
// }

// // Display list of 10 things in json
// func exerciseTwo() {
// 	todos := Todos{}
// 	todos.AddTaskItems(tasksToAdd...)
// 	fmt.Println(todos.Json())
// }

// func exerciseThree() {
// 	todos := Todos{}
// 	todos.AddTaskItems(tasksToAdd...)
// 	outputFile, _ := os.Create("./output.json")
// 	todos.Save(outputFile)
// }

// // read a todo list from a json file
// func exerciseFour() {
// 	fs := os.DirFS(".")
// 	newTodos := Todos{}
// 	newTodos.Load(fs, "todos.json")
// 	newTodos.PrintDescriptions(os.Stdout)
// }
