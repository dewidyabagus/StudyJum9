package main

import "fmt"

/*
	Description : Variadic function adalah suatu function yang memiliki parameter dengan type yang sama dan tak terbatas
	Declaration :
	func namaFunction(params ...dataType) {
		// params = []dataType
	}
	func main() {
		nameFunction(1)
		nameFunction(1,2)
		nameFunction(1,2,3)
	}

	Invalid function:
	func nameFunc(params ...dataType, max int64) {

	}

	Valid function:
	func nameFunc(name string, max int64, params ...dataType) {

	}
*/

func main() {
	var result = sum(2, 4, 5, 7, 8)
	fmt.Println("Sum 2 + 4 + 5 + 7 + 8 :", result)
}

func sum(numbers ...uint64) (result uint64) {
	for _, num := range numbers {
		result += num
	}

	return
}
