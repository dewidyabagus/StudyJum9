package main

import "fmt"

func main() {
	// Suatu proses yang digunakan untuk memasukan proses eksekusi ke dalam antrian dengan model LIFO (Last In First Out)

	defer fmt.Println("Cetak defer 1")
	defer fmt.Println("Cetak defer 2")
	defer fmt.Println("Cetak defer 3")

	fmt.Println("Cetak baris 12")
	fmt.Println("Cetak baris 13")
	fmt.Println("Cetak baris 14")
	/*
		Result:
			Cetak baris 12
			Cetak baris 13
			Cetak baris 14
			Cetak defer 3
			Cetak defer 2
			Cetak defer 1
	*/
}
