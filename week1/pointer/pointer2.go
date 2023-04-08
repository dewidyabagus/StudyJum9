package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println("#1 nilai x =", x)

	// simbol & untuk proses referencing
	// simbol * untuk proses dereferencing
	var y *int = &x
	fmt.Println("#1 nilai y =", y)   // Result alamat memory x
	fmt.Println("#1 nilai *y =", *y) // Result nilai pada alamat memory x

	var z int = *y
	fmt.Println("#1 nilai z =", z)
}
