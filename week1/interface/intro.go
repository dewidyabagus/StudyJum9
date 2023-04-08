package main

import "fmt"

func printType(value interface{}) {
	// Harus berhati2 ketika menentukan data type dari interface{}
	// var result = 1 + value.(int64)
	// fmt.Println("Result:", result)

	switch val := value.(type) {
	case []int:
		fmt.Println("TO DO: Action this value []int:", val)

	case []string:
		fmt.Println("TO DO: Action this value []string:", val)

	case bool:
		fmt.Println("TO DO: Action this value boolean:", val)
	}
}

func main() {
	var numbers = []int{1}
	printType(numbers)

	var data = []string{"A"}
	printType(data)

	var result = true
	printType(result)
}
