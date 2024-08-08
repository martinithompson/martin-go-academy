// ******* Step 1: Makes a race condition via go routine access to a shared resource ******

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	input := []int{1, 2, 3, 4, 5}
// 	result := []int{} // shared resource

// 	for _, data := range input {
// 		wg.Add(1)
// 		go processData(&wg, &result, data)
// 	}

// 	wg.Wait()
// 	fmt.Println(result)
// }

// func processData(wg *sync.WaitGroup, res *[]int, data int) {
// 	defer wg.Done()

// 	*res = append(*res, data*2) //dereference result
// }

// ******* Step 2: Implement mutex ******
// ******* but not ideal as causes bottleneck, each goroutine must wait for fn to be unlocked ******
// ******* so basically just running synchronously anyway, takes 10s to run ******

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var lock sync.Mutex

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	input := []int{1, 2, 3, 4, 5}
// 	result := []int{} // shared resource

// 	for _, data := range input {
// 		wg.Add(1)
// 		go processData(&wg, &result, data)
// 	}

// 	wg.Wait()
// 	fmt.Println(result)
// 	fmt.Println(time.Since(start))
// }

// func process(data int) int {
// 	time.Sleep(time.Second * 2)
// 	return data * 2
// }

// func processData(wg *sync.WaitGroup, res *[]int, data int) {
// 	lock.Lock()
// 	defer wg.Done()

// 	processedData := process(data)
// 	*res = append(*res, processedData) //dereference result
// 	lock.Unlock()
// }

// ******* Step 3: Improve mutex ******
// ******* lock only the critical shared write access ******
// ******* better as time consuming step can happen concurrently, takes 2s to run ******
// ******* but if critical step is time consuming then mutex may still not be ideal ******

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var lock sync.Mutex

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	input := []int{1, 2, 3, 4, 5}
// 	result := []int{} // shared resource

// 	for _, data := range input {
// 		wg.Add(1)
// 		go processData(&wg, &result, data)
// 	}

// 	wg.Wait()
// 	fmt.Println(result)
// 	fmt.Println(time.Since(start))
// }

// func process(data int) int {
// 	time.Sleep(time.Second * 2)
// 	return data * 2
// }

// func processData(wg *sync.WaitGroup, res *[]int, data int) {
// 	defer wg.Done()
// 	processedData := process(data)

// 	lock.Lock()
// 	*res = append(*res, processedData) //dereference result
// 	lock.Unlock()
// }

// ******* Step 4: Confinement ******
// ******* update an individual element of the shared data store (slice) so there is no actual conflict ******
// ******* runs in sequence, takes 2s ******

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

func processData(wg *sync.WaitGroup, resDestination *int, data int) {
	defer wg.Done()
	processedData := process(data)

	*resDestination = processedData
}
