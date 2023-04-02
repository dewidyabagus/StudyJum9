package main

import "fmt"

func main() {
labelMainFor:
	for i := 0; i < 5; i++ {
	labelForSub1:
		for n := i; n >= 0; n-- {
			if i == 3 && n == 2 {
				break labelMainFor
			}
			fmt.Printf("Perulangan i = %d, pada subfor n = %d \n", i, n)
			for m := n; m > 0; m-- {
				continue labelForSub1
				// fmt.Printf("Perulangan i = %d, pada subfor n = %d \n", i, n)
			}
		}
		fmt.Println()
	}
	fmt.Println("Kembali ke blok function")
}
