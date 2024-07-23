package main

import (
	"fmt"
	"math"
)

func numberOfDigits(number int) int {
	return int(math.Log10(float64(number))) + 1
}

func readInts(size int) [3]int {
	var arr [3]int
	var readInt int

	for i := 0; i < 3; i++ {
		needed := true
		for needed {
			fmt.Printf("Please enter a %d digit number", size)
			_, err := fmt.Scanf("%d", &readInt)
			if err == nil && numberOfDigits(readInt) == size {
				arr[i] = readInt
				needed = false
			}
		}
	}
	return arr
}

func sumArray(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	singleDigitArr := readInts(1)
	fmt.Printf("%v", singleDigitArr)

	twoDigitArr := readInts(2)
	fmt.Printf("%v", twoDigitArr)

	threeDigitArr := readInts(3)
	fmt.Printf("%v", threeDigitArr)

	sumAllArrays := sumArray(singleDigitArr) + sumArray(twoDigitArr) + sumArray(threeDigitArr)

	fmt.Println("The sum of all elements is", sumAllArrays)
}
