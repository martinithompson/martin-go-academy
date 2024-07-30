package main

import (
	"fmt"
	"sync"
)

func printTodoItems(wg *sync.WaitGroup, ch chan string, items []string) {
	defer wg.Done()
	for _, item := range items {
		ch <- item
	}
	close(ch)
}

func printTodoStatuses(wg *sync.WaitGroup, ch chan string, statuses []string) {
	defer wg.Done()
	for _, status := range statuses {
		ch <- status
	}
	close(ch)
}

func main() {
	items := []string{
		"Wash the car",
		"Go shopping",
		"Phone the bank",
		"Clean the house",
		"Finish report",
		"Book appointment",
		"Exercise",
		"Make dinner",
		"Read a book",
		"Pay bills",
	}
	statuses := []string{
		"Pending",
		"In Progress",
		"Completed",
		"Pending",
		"In Progress",
		"Completed",
		"Pending",
		"In Progress",
		"Completed",
		"Pending",
	}

	itemChannel := make(chan string)
	statusChannel := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go printTodoItems(&wg, itemChannel, items)
	go printTodoStatuses(&wg, statusChannel, statuses)

	for i := 0; i < len(items); i++ {
		item := <-itemChannel
		status := <-statusChannel
		fmt.Printf("Item: %s, Status: %s\n", item, status)
	}

	wg.Wait()
}
