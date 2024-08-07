package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todo-app/project/todos"
)

const (
	apiBaseUrl = "http://localhost:8000/todos"
	port       = ":8080"
)

var renderer, _ = NewTodoRenderer()
var localStore = todos.TodoList{}

func handleRenderList(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(apiBaseUrl)
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
	localStore.Items = todoData
	err = renderer.List(w, localStore.Items)
	handleHttpError(w, err)
}

func handleRenderAdd(w http.ResponseWriter, r *http.Request) {
	err := renderer.Add(w, nil)
	handleHttpError(w, err)
}

func handleAddTodo(writer http.ResponseWriter, request *http.Request) {
	item, _ := getTodoFromForm(request)
	newTodo := todos.Todo{Name: item, Completed: false}
	jsonData := convertTodoToJSON(newTodo)

	resp, err := http.Post(apiBaseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	validateStatusCode(resp, http.StatusCreated)
	navigateToListPage(writer, request)
}

func main() {
	http.HandleFunc("/list", handleRenderList)
	http.HandleFunc("/add", handleRenderAdd)
	http.HandleFunc("/add-todo", handleAddTodo)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleHttpError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}

// func errorCheck(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func validateStatusCode(resp *http.Response, expectedStatusCode int) {
	if resp.StatusCode != expectedStatusCode {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}
}

func navigateToListPage(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/list", http.StatusFound)
}

func getTodoFromForm(request *http.Request) (string, bool) {
	item := request.FormValue("item")
	completed := request.FormValue("completed")

	return item, checkboxValueToBool(completed)
}

func convertTodoToJSON(todo todos.Todo) []byte {
	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	return jsonData
}

// func getIdFromPath(path string) string {
// 	pathParams := strings.Split(path, "/")
// 	return pathParams[len(pathParams)-1]
// }

func checkboxValueToBool(value string) bool {
	return value == "on"
}

// func getTodoById(id string) todos.Todo {
// 	for _, todo := range localStore.Items {
// 		if todo.Id == id {
// 			return todo
// 		}
// 	}
// 	return todos.Todo{}
// }

// func listPageHandler(writer http.ResponseWriter, request *http.Request) {
// 	resp, err := http.Get(apiBaseUrl)
// 	if err != nil {
// 		log.Fatalf("Error making GET request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	validateStatusCode(resp, http.StatusOK)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}

// 	var todoData []todos.Todo
// 	err = json.Unmarshal(body, &todoData)
// 	if err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	// save to global state
// 	localStore.Items = todoData
// 	// renderTemplate(writer, "list.html", todoData)
// }

// func addPageHandler(writer http.ResponseWriter, request *http.Request) {
// 	tmpl, err := template.ParseFiles("add.html")
// 	errorCheck(err)
// 	err = tmpl.Execute(writer, nil)
// 	errorCheck(err)
// }

// func addTodoHandler(writer http.ResponseWriter, request *http.Request) {
// 	item, _ := getTodoFromForm(request)
// 	newTodo := todos.Todo{Name: item, Completed: false}
// 	jsonData := convertTodoToJSON(newTodo)

// 	resp, err := http.Post(apiBaseUrl, "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		log.Fatalf("Error making POST request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	validateStatusCode(resp, http.StatusCreated)
// 	navigateToListPage(writer, request)
// }

// func editPageHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := getIdFromPath(request.URL.Path)
// 	selectedTodo := getTodoById(id)
// 	// renderTemplate(writer, "edit.html", selectedTodo)
// }

// func editTodoHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := getIdFromPath(request.URL.Path)
// 	item, completed := getTodoFromForm(request)
// 	updatedTodo := todos.Todo{Id: id, Name: item, Completed: completed}
// 	jsonData := convertTodoToJSON(updatedTodo)

// 	patchUrl := fmt.Sprintf("%s/%s", apiBaseUrl, id)

// 	req, err := http.NewRequest(http.MethodPatch, patchUrl, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		log.Fatalf("Error creating PATCH request: %v", err)
// 	}
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending PATCH request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	validateStatusCode(resp, http.StatusCreated)
// 	navigateToListPage(writer, request)
// }

// func deletePageHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := getIdFromPath(request.URL.Path)
// 	selectedTodo := getTodoById(id)
// 	renderTemplate(writer, "delete.html", selectedTodo)
// }

// func deleteTodoHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := getIdFromPath(request.URL.Path)
// 	deleteUrl := fmt.Sprintf("%s/%s", apiBaseUrl, id)

// 	req, err := http.NewRequest(http.MethodDelete, deleteUrl, nil)
// 	if err != nil {
// 		log.Fatalf("Error creating DELETE request: %v", err)
// 	}
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending DELETE request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	validateStatusCode(resp, http.StatusNoContent)
// 	navigateToListPage(writer, request)
// }
