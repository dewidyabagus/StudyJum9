package main

import "time"

/*
	Latar belakang generic function. Mulai dikenalkan di go versi 1.18 dan terbaru

	func sum(numbers []int)            {}
	func sumFloat64(numbers []float64) {}

	func main() {
		var numbersInt = []int{1, 2, 3, 4}
		sum(numbersInt)						// Valid

		var numbersFloat64 = []float64{1, 2, 3, 4}
		sum(numbersFloat64)					// Invalid
		sumFloat64(numbersFloat64)			// Valid
	}
*/

type collection interface {
	string | bool | int | time.Time
}

func sum[T []int | []float64](numbers T) {}

// func genericExample[T string | bool | int | time.Time](value T) {} | versi panjang

func genericExample[T collection](value T) {}

// func genericExample2[T collection](value T) {}
// func genericExample3[T collection](value T) {}

func main() {
	var numbersInt = []int{1, 2, 3, 4}
	sum(numbersInt)

	var numbersFloat64 = []float64{1, 2, 3, 4}
	sum(numbersFloat64)
}
