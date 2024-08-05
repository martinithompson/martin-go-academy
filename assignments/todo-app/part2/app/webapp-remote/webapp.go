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

const url = "http://localhost:5000/todos"

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type TodosData struct {
	Todo  todos.Todo
	Index int
}

func listPageHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("list.html")

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

	var todoData []todos.Todo
	err = json.Unmarshal(body, &todoData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// save to global state
	todoList.Items = todoData

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

	todo := todos.Todo{
		Item:      item,
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

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

	todo := todos.Todo{
		Item:      item,
		Completed: completed == "on",
	}

	fmt.Println("todo: ", todo)

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	patchUrl := fmt.Sprintf("%s/%d", url, index)

	req, err := http.NewRequest(http.MethodPatch, patchUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating PATCH request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending PATCH request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	fmt.Println("Resource updated successfully")

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func deletePageHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	fmt.Println("Index: ", todoList.Items[index])
	tmpl, err := template.ParseFiles("delete.html")
	errorCheck(err)
	data := TodosData{Todo: todoList.Items[index], Index: index}
	err = tmpl.Execute(writer, data)
	errorCheck(err)
}

func deleteTodoHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)

	deleteUrl := fmt.Sprintf("%s/%d", url, index)

	fmt.Println("***deleteURL***", deleteUrl)

	req, err := http.NewRequest(http.MethodDelete, deleteUrl, nil)
	if err != nil {
		log.Fatalf("Error creating DELETE request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	fmt.Println("Resource deleted successfully")

	http.Redirect(writer, request, "/list", http.StatusFound)
}

func getIdFromPath(path string) string {
	pathParams := strings.Split(path, "/")
	return pathParams[len(pathParams)-1]
}

func checkboxValueToBool(value string) bool {
	return value == "on"
}

const port = ":8080"

func main() {
	http.HandleFunc("/list", listPageHandler)
	http.HandleFunc("/add", addPageHandler)
	http.HandleFunc("/add-todo", addTodoHandler)
	http.HandleFunc("/edit/", editPageHandler)
	http.HandleFunc("/edit-todo/", editTodoHandler)
	http.HandleFunc("/delete/", deletePageHandler)
	http.HandleFunc("/delete-todo/", deleteTodoHandler)

	log.Printf("Starting webapp on http://localhost%s", port)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
