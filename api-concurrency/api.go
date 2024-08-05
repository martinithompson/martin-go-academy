package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CounterStore struct {
	counters map[string]int
}

func (cs CounterStore) get(w http.ResponseWriter, req *http.Request) {
	log.Printf("get %v", req)
	name := req.URL.Query().Get("name")
	if val, ok := cs.counters[name]; ok {
		fmt.Fprintf(w, "%s: %d\n", name, val)
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

func (cs CounterStore) set(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	name := req.URL.Query().Get("name")
	val := req.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		cs.counters[name] = intval
		fmt.Fprintf(w, "ok\n")
	}
}

func main() {
	store := CounterStore{counters: map[string]int{"i": 0, "j": 0}}
	http.HandleFunc("/get", store.get)
	http.HandleFunc("/set", store.set)
	http.HandleFunc("/inc", store.inc)

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
	SetCommand
	IncCommand
)

type Command struct {
	ty        CommandType
	name      string
	val       int
	replyChan chan int
}

func startCounterManager(initvals map[string]int) chan<- Command {
	counters := make(map[string]int)
	for k, v := range initvals {
		counters[k] = v
	}

	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			switch cmd.ty {
			case GetCommand:
				if val, ok := counters[cmd.name]; ok {
					cmd.replyChan <- val
				} else {
					cmd.replyChan <- -1
				}
			case SetCommand:
				counters[cmd.name] = cmd.val
				cmd.replyChan <- cmd.val
			case IncCommand:
				if _, ok := counters[cmd.name]; ok {
					counters[cmd.name]++
					cmd.replyChan <- counters[cmd.name]
				} else {
					cmd.replyChan <- -1
				}
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

// func (cs CounterStore) inc(w http.ResponseWriter, req *http.Request) {
// 	log.Printf("inc %v", req)
// 	name := req.URL.Query().Get("name")
// 	if _, ok := cs.counters[name]; ok {
// 		cs.counters[name]++
// 		fmt.Fprintf(w, "ok\n")
// 	} else {
// 		fmt.Fprintf(w, "%s not found\n", name)
// 	}
// }

func (s *Server) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	name := req.URL.Query().Get("name")
	replyChan := make(chan int)
	s.cmds <- Command{ty: IncCommand, name: name, replyChan: replyChan}

	reply := <-replyChan
	if reply >= 0 {
		fmt.Fprintf(w, "ok\n")
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}
