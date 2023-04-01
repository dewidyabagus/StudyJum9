package main

import "fmt"

func main() {
	/*
		Comment Multi Line
		Input Comment In Here
	*/

	// Long declaration var without initializer
	var roomNumber int
	roomNumber = 135
	fmt.Println("Room Number:", roomNumber)

	// Long declaration var with initializer
	var floorNumber int = 13
	fmt.Println("Floor Number:", floorNumber)

	// Short declaration
	firstName, middleName, lastName := "Johns", "Aliando", "Wick"
	fmt.Println("Full Name:", firstName, middleName, lastName)

	// Typed Constants
	const pi float64 = 3.14
	fmt.Println("PI:", pi)

	// Untype Constant
	const zipCode = 11111
	fmt.Println("Zip Code:", zipCode)
}
