package main

import "fmt"

func main() {
	// Tipe data array masuk ke aggregate data type
	// var (identifier) [(length)](data-type) | Versi panjang
	var names [4]string
	names[0] = "Agus"
	names[1] = "Achmad"
	names[2] = "Eko"
	names[3] = "Bagus"
	fmt.Println("Names is", names)

	// Versi singkat
	// var names = [4]string{"Agus", "Achmad", "Eko", "Bagus"}
}
