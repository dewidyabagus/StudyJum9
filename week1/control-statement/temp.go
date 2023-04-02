package main

import (
	"fmt"
	"math"
)

func main() {
	var edge float64 = 20
	// var name string | Scope variable di semua blok dalam fucntion

	// cubeVolume variable ini scope hanya pada block if - else if - else
	if cubeVolume := math.Pow(edge, 3); cubeVolume > 10000 {
		// fmt.Println(name) | Valid
		fmt.Println("Lerge volumne cube. Volume:", cubeVolume)

	} else if cubeVolume >= 5000 {
		// fmt.Println(name) | Valid
		fmt.Println("Medium volume cube. Volume:", cubeVolume)

	} else {
		fmt.Println("Small volume cube. Volume:", cubeVolume)
	}

	// fmt.Println(cubeVolume) | Tidak valid
}
