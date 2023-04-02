package main

import "fmt"

func main() {
	// Penggunaan keyword for - range
	/*
		for (index/ key), (value) := range (array, slice, map, string) {
			// blok kode
		}
	*/
	// var names = [4]string{"budi", "eko", "rachmad", "agus"}
	// for i, name := range names {
	// 	fmt.Printf("Index: %d, Name: %s \n", i, name)
	// }

	// String adalah kumpulan dari tipe data rune
	var name = "Nest Academy"
	for i, val := range name {
		fmt.Printf("Index: %d, asci: %v, val: %s \n", i, val, string(val))
	}
}
