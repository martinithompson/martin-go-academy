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

func (i *InMemoryTodoStore) DeleteTodo(index int) {
	if index < 0 || index >= len(i.store.Items) {
		return
	}
	i.store.DeleteTodoItem(index)
}

func (i *InMemoryTodoStore) EditTodo(index int, updated todos.Todo) {
	if index < 0 || index > len(i.store.Items) {
		return
	}
	if i.store.Items[index-1].Item != updated.Item {
		i.store.UpdateTodoItem(index, updated.Item)
	}
	if i.store.Items[index-1].Completed != updated.Completed {
		i.store.ToggleTodoCompleted(index)
	}
}

func (i *InMemoryTodoStore) ToggleTodoCompleted(index int) {
	i.store.ToggleTodoCompleted(index)
}

func (i *InMemoryTodoStore) GetTodos() []todos.Todo {
	return i.store.Items
}
