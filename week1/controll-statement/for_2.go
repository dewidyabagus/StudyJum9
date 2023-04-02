package main

import "fmt"

func main() {
	/*
		Perulangan for seperti while

		1. Declaration variable
		2. Keyword for dengan kondisi
			for (kondisi) {
				// blok kode
				// increment / decrement
			}
	*/
	var i int
	for i < 5 {
		fmt.Println("#1 Perulangan dengan nilai i:", i)
		i++
	}

	var n int
	for ; n < 5; n++ {
		fmt.Println("#2 Perulangan dengan nilai i:", n)
	}
}
