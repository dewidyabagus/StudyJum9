package main

import "fmt"

func main() {
	// For secara umum
	/*
		for (inisialisasi); (kondisi); (increment/decrement) {
			// blok kode
		}

	*/
	/*
		var i int
		for i = 0; i < 5; i++ {
			fmt.Println("Perulangan dengan nilai i:", i)
		}
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("#1 Perulangan dengan nilai i:", i)
	}

	fmt.Println()

	for i := 0; i < 5; i += 2 {
		fmt.Println("#2 Perulangan dengan nilai i:", i)
	}
	// fmt.Println(i) | Invalid
}
