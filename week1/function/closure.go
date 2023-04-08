package main

import "fmt"

/*
	closure adalah suatu fungsi yang disimpan dalam variable (anonymous function)

	declaration :
	func main() {
		// Fungsi yang disimpan dalam bentuk variable
		var getMaxNumbers = func(numbers []int) int {
		}

		// IIFE (Immediately-Invoked Function Expression)
		var maxNumbers = func(numbers []int) int {
		}(numbers)
	}

	// Fungsi juga bisa mengembalikan suatu fungsi
	func findMax(numbers []int, max int) (int, func() []int) {
		return 0, func() []int {
			return []int{}
		}
	}
*/

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var getMaxNumbers = func(numbers []int) int {
		var max int

		for _, val := range numbers {
			if val > max {
				max = val
			}
		}

		return max
	}

	fmt.Println("Fungsi disimpan dalam variable:", getMaxNumbers(numbers))

	var maxNumber = func(numbers []int) int {
		var max int

		for _, val := range numbers {
			if val > max {
				max = val
			}
		}

		return max
	}(numbers)

	fmt.Println("Fungsi IIFE:", maxNumber)

	var length, details = findMax(numbers, 2)
	fmt.Println("Length:", length)
	fmt.Println("Details:", details())
}

func findMax(numbers []int, max int) (int, func() []int) {
	// var result []int | kurang effective
	var result = make([]int, 0, len(numbers))

	for _, num := range numbers {
		if num <= max {
			result = append(result, num)
		}
	}

	return len(result), func() []int {
		return result
	}
}
