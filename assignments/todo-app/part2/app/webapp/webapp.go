package main

import (
	"html/template"
	"log"
	"net/http"
	"todo-app/project/todos"
)

var todoList = todos.Todos{}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func sayHelloHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello world")
}

func viewTodosHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("list.html")
	tmpl.Execute(writer, todoList.Items)
}

func addTodoHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("add.html")
	errorCheck(err)
	err = tmpl.Execute(writer, nil)
	errorCheck(err)
}

func createTodoHandler(writer http.ResponseWriter, request *http.Request) {
	item := request.FormValue("item")
	todoList.AddTaskItems(item)
	http.Redirect(writer, request, "/list", http.StatusFound)
}

func toggleTodoHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, request.URL.Path)
}

func main() {
	// todoList.AddTaskItems("hello world")
	http.HandleFunc("/hello", sayHelloHandler)
	http.HandleFunc("/list", viewTodosHandler)
	http.HandleFunc("/add", addTodoHandler)
	http.HandleFunc("/create", createTodoHandler)
	http.HandleFunc("/toggle/", toggleTodoHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
