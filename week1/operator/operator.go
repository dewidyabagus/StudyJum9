package main

import "fmt"

func main() {
	var discont float64 = 0.1 * 10000
	fmt.Println("Total discont:", discont)
	// Result: Total discont: 1000

	var left, right = 13, 13
	fmt.Println("Balanced:", (left == right))
	// Result: Balanced: true

	var total = 10000
	fmt.Println("Voucher:", (total >= 1000) && (total <= 10000))
	// Result: Voucher: true

	// total += 1000
	// Sama dengan: total = total + 1000

	// total *= 2
	// Sama dengan: total = total * 2

	// Auti Incriment / Decriment
	// i++ / i = i + 1
	// i-- / i = i - 1

	var opr uint8 = 255
	fmt.Println("Total:", opr+1) // Expect: 256, actual: 0
}
