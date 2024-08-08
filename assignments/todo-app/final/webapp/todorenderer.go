package main

import (
	"embed"
	"html/template"
	"io"
	"todo-app/project/todos"
)

var (
	//go:embed "templates/*"
	todoTemplates embed.FS
)

type TodoRenderer struct {
	templ *template.Template
}

func NewTodoRenderer() (*TodoRenderer, error) {
	templ, err := template.ParseFS(todoTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &TodoRenderer{templ: templ}, nil
}

func (r *TodoRenderer) List(w io.Writer, t []todos.Todo) error {
	return executeTemplate(r, w, "list.gohtml", t)
}

func (r *TodoRenderer) Add(w io.Writer) error {
	// todo - generic executeTemplate function wont accept nil for T parameter??
	if err := r.templ.ExecuteTemplate(w, "add.gohtml", nil); err != nil {
		return err
	}

	return nil
}

func (r *TodoRenderer) Edit(w io.Writer, t todos.Todo) error {
	return executeTemplate(r, w, "edit.gohtml", t)
}

func (r *TodoRenderer) Delete(w io.Writer, t todos.Todo) error {
	return executeTemplate(r, w, "delete.gohtml", t)
}

func executeTemplate[T todos.Todo | []todos.Todo](r *TodoRenderer, w io.Writer, f string, d T) error {
	if err := r.templ.ExecuteTemplate(w, f, d); err != nil {
		return err
	}

	return nil
}
