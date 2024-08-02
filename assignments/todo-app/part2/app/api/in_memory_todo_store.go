package main

import (
	todos "todo-app/project/todos"
)

func NewInMemoryTodoStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{todos.Todos{}}
}

type InMemoryTodoStore struct {
	store todos.Todos
}

func (i *InMemoryTodoStore) AddTodo(todo todos.Todo) {
	i.store.AddTodoItems(todo)
}

func (i *InMemoryTodoStore) GetTodos() []todos.Todo {
	return i.store.Items
}
