package main

import (
	"log"
	todos "todo-app/project/todos"
)

type CommandType int

const (
	GetCommand CommandType = iota
	AddCommand
	EditCommand
	DeleteCommand
)

type Command struct {
	ct        CommandType
	todoItem  todos.Todo
	replyChan chan []todos.Todo
}

type Store struct {
	TodoList todos.TodoList
}

func startTodoManager() chan<- Command {
	s := Store{}

	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			switch cmd.ct {
			case GetCommand:
				cmd.replyChan <- s.TodoList.Items
			case AddCommand:
				s.TodoList.AddTodosByName(cmd.todoItem.Name)
				cmd.replyChan <- s.TodoList.Items
			case EditCommand:
				s.TodoList.UpdateTodo(cmd.todoItem)
				cmd.replyChan <- s.TodoList.Items
			case DeleteCommand:
				s.TodoList.DeleteTodo(cmd.todoItem.Id)
				cmd.replyChan <- s.TodoList.Items
			default:
				log.Fatal("unknown command", cmd.ct)
			}
		}
	}()
	return cmds
}
