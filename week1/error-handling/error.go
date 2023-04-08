package main

import "fmt"

func printType(value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Hasil recovery:", r)
		}
	}()

	var total = value.(int) + 1

	fmt.Println("Total:", total)
}

func main() {
	printType(uint64(10))

	fmt.Println("Baris 15")
	fmt.Println("Baris 16")
	fmt.Println("Baris 17")
}
