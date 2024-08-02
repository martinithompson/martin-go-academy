package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo-app/project/todos"
)

var todoList = todos.Todos{}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type TodoData struct {
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type TodosData struct {
	Todo  todos.Todo
	Index int
}

func listPageHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("list.html")

	url := "http://localhost:5000/todos"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println(string(body))

	var todoData []TodoData
	err = json.Unmarshal(body, &todoData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	tmpl.Execute(writer, todoData)
}

func addPageHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("add.html")
	errorCheck(err)
	err = tmpl.Execute(writer, nil)
	errorCheck(err)
}

func addTodoHandler(writer http.ResponseWriter, request *http.Request) {
	item := request.FormValue("item")
	// todoList.AddTaskItems(item)

	todo := TodoData{
		Item:      item,
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	url := "http://localhost:5000/todos"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func editPageHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	tmpl, err := template.ParseFiles("edit.html")
	errorCheck(err)
	data := TodosData{Todo: todoList.Items[index], Index: index}
	err = tmpl.Execute(writer, data)
	errorCheck(err)
}

func editTodoHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	item := request.FormValue("item")
	completed := request.FormValue("completed")
	if todoList.Items[index].Item != item {
		todoList.UpdateTodoItem(index+1, item)
	}
	if todoList.Items[index].Completed != checkboxValueToBool(completed) {
		todoList.ToggleTodoCompleted(index + 1)
	}

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func deletePageHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	tmpl, err := template.ParseFiles("delete.html")
	errorCheck(err)
	data := TodosData{Todo: todoList.Items[index], Index: index}
	err = tmpl.Execute(writer, data)
	errorCheck(err)
}

func deleteTodoHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	todoList.DeleteTodoItem(index + 1)

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func getIdFromPath(path string) string {
	pathParams := strings.Split(path, "/")
	return pathParams[len(pathParams)-1]
}

func checkboxValueToBool(value string) bool {
	return value == "on"
}

func main() {
	http.HandleFunc("/list", listPageHandler)
	http.HandleFunc("/add", addPageHandler)
	http.HandleFunc("/add-todo", addTodoHandler)
	http.HandleFunc("/edit/", editPageHandler)
	http.HandleFunc("/edit-todo/", editTodoHandler)
	http.HandleFunc("/delete/", deletePageHandler)
	http.HandleFunc("/delete-todo/", deleteTodoHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
