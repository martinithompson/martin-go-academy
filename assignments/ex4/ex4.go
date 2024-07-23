package main

import "fmt"

func populateArray() [10]int {
	var numberArray [10]int

	for i := 1; i <= 10; i++ {
		numberArray[i-1] = i
	}

	return numberArray
}

func displayArray(arr [10]int, ascending bool) {
	var evens []int
	var odds []int
	if ascending {
		for i := 0; i < len(arr); i++ {
			evens, odds = appendEvenOrOdd(evens, odds, arr[i])
			fmt.Println(arr[i])
		}
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			evens, odds = appendEvenOrOdd(evens, odds, arr[i])
			fmt.Println(arr[i])
		}
	}

	fmt.Printf("The even numbers are %v and the odd nuumbers are %v\n", evens, odds)
}

func appendEvenOrOdd(evens, odds []int, value int) ([]int, []int) {
	if isEven(value) {
		evens = append(evens, value)
	} else {
		odds = append(odds, value)
	}
	return evens, odds
}

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	myArray := populateArray()

	displayArray(myArray, true)

	displayArray(myArray, false)
}
