package main

import (
	"fmt"
	"strings"
)

type CondString func(string) bool
type NewString string

func filter(data []string, condition CondString) []string {
	var result = make([]string, 0, len(data))

	for _, val := range data {
		if ok := condition(val); ok {
			result = append(result, val)
		}
	}

	return result
}

/*
	func filterText1(data []string, condition func(string)bool) {}
	func filterText2(data []string, condition func(string)bool) {}
*/

func main() {
	var data = []string{"hello", "world", "john", "heru", "wick"}

	var filterContainsO = filter(data, func(text string) bool {
		return strings.Contains(text, "o")
	})
	fmt.Println("Result filterContainsO :", filterContainsO)

	var filterLength5 = filter(data, func(text string) bool {
		return len(text) == 5
	})
	fmt.Println("Result filterLength5 :", filterLength5)
}
