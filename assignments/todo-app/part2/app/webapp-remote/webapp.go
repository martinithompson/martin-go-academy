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

const (
	url  = "http://localhost:5000/todos"
	port = ":8080"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type TodosData struct {
	Todo  todos.Todo
	Index int
}

func validateStatusCode(resp *http.Response, expectedStatusCode int) {
	if resp.StatusCode != expectedStatusCode {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}
}

func renderTemplate(writer http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles(templateName)
	errorCheck(err)
	err = tmpl.Execute(writer, data)
	errorCheck(err)
}

func navigateToListPage(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/list", http.StatusFound)
}

func getTodoIndex(path string) int {
	index, err := strconv.Atoi(getIdFromPath(path))
	errorCheck(err)
	return index
}

func getTodoFromForm(request *http.Request) todos.Todo {
	item := request.FormValue("item")
	completed := request.FormValue("completed")

	return todos.Todo{
		Item:      item,
		Completed: checkboxValueToBool(completed),
	}
}

func convertTodoToJSON(todo todos.Todo) []byte {
	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	return jsonData
}

func getIdFromPath(path string) string {
	pathParams := strings.Split(path, "/")
	return pathParams[len(pathParams)-1]
}

func checkboxValueToBool(value string) bool {
	return value == "on"
}

func listPageHandler(writer http.ResponseWriter, request *http.Request) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	validateStatusCode(resp, http.StatusOK)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var todoData []todos.Todo
	err = json.Unmarshal(body, &todoData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// save to global state
	todoList.Items = todoData
	renderTemplate(writer, "list.html", todoData)
}

func addPageHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("add.html")
	errorCheck(err)
	err = tmpl.Execute(writer, nil)
	errorCheck(err)
}

func addTodoHandler(writer http.ResponseWriter, request *http.Request) {
	todo := getTodoFromForm(request)
	jsonData := convertTodoToJSON(todo)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	validateStatusCode(resp, http.StatusCreated)
	navigateToListPage(writer, request)
}

func editPageHandler(writer http.ResponseWriter, request *http.Request) {
	index, err := strconv.Atoi(getIdFromPath(request.URL.Path))
	errorCheck(err)
	data := TodosData{Todo: todoList.Items[index], Index: index}
	renderTemplate(writer, "edit.html", data)
}

func editTodoHandler(writer http.ResponseWriter, request *http.Request) {
	index := getTodoIndex(request.URL.Path)
	todo := getTodoFromForm(request)
	jsonData := convertTodoToJSON(todo)

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

	validateStatusCode(resp, http.StatusCreated)
	navigateToListPage(writer, request)
}

func deletePageHandler(writer http.ResponseWriter, request *http.Request) {
	index := getTodoIndex(request.URL.Path)
	data := TodosData{Todo: todoList.Items[index], Index: index}
	renderTemplate(writer, "delete.html", data)
}

func deleteTodoHandler(writer http.ResponseWriter, request *http.Request) {
	index := getTodoIndex(request.URL.Path)
	deleteUrl := fmt.Sprintf("%s/%d", url, index)

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

	validateStatusCode(resp, http.StatusNoContent)
	navigateToListPage(writer, request)
}

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
