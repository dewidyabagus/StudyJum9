package main

import (
	"errors"
	"fmt"
)

// func init() {
// 	fmt.Println("Function init")
// }

func main() {
	printHello("Widya")

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang(5, 6))

	if res, err := pembagian(6, 0); err == nil {
		fmt.Println("pembagian 6 / 0:", res)
	} else {
		fmt.Println("Error:", err.Error())
	}

}

// Function tanpa nilai kembali
func printHello(name string) {
	fmt.Println("Hello, name", name)
}

// Function dengan nilai kembali
func luasPersegiPanjang(p, l int64) int64 {
	return p * l
}

// Function dengan multiple retrun value
func pembagian(p, l int64) (int64, error) {
	if l == 0 {
		return 0, errors.New("nilai l adalah 0, tidak dapat diproses")
	}
	return p / l, nil
}
