package main

import "fmt"

func main() {
	// Keyword break dalam for
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("#1 Perulangan dengan nilai i:", i)
	}
	fmt.Println("Melanjutkan squencial code diluar for")

	// Keyword continue dalam for
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println("#2 Perulangan dengan nilai i:", i)
	}
	fmt.Println("Melanjutkan squencial code diluar for (for keyword continue)")
}
