package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type TodoStore struct {
	todos []Todo
}

func main() {
	server := Server{}
	server.cmds = startTodoManager()

	http.HandleFunc("/inc", server.inc)
	http.HandleFunc("/get", server.getTodos)

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}

type CommandType int

const (
	GetCommand CommandType = iota
	AddCommand
)

type Command struct {
	ty        CommandType
	name      string
	replyChan chan []Todo
}

func startTodoManager() chan<- Command {
	todos := TodoStore{}

	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			switch cmd.ty {
			case GetCommand:
				fmt.Println("GetCommand")
				cmd.replyChan <- todos.todos
			case AddCommand:
				fmt.Println("AddCommand")
				todos.todos = append(todos.todos, Todo{Name: cmd.name, Completed: false})
				cmd.replyChan <- todos.todos
			default:
				log.Fatal("unknown command type", cmd.ty)
			}
		}
	}()
	return cmds
}

type Server struct {
	cmds chan<- Command
}

func (s *Server) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	name := req.URL.Query().Get("name")
	replyChan := make(chan []Todo)
	s.cmds <- Command{ty: AddCommand, name: name, replyChan: replyChan}

	reply := <-replyChan
	if reply != nil {
		fmt.Fprintf(w, "ok\n")
	}
}

func (s *Server) getTodos(w http.ResponseWriter, req *http.Request) {
	log.Printf("get %v", req)
	replyChan := make(chan []Todo)
	s.cmds <- Command{ty: GetCommand, replyChan: replyChan}

	todos := <-replyChan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
