package main

import "fmt"

func main() {
	// var names = []string{"Agus", "Achamd", "Eko", "Bagus"}
	// fmt.Println("Length:", len(names))
	// fmt.Println("Capacity:", cap(names))

	// Result:
	// Length => 4
	// Capacity => 4
	/*
		fmt.Println("Names names:", names)

		names2 := names
		fmt.Printf("Memory names: %p \n", names)
		fmt.Printf("Memory names2: %p \n", names)

		names[0] = "Juki"

		fmt.Println("Names :", names)
		fmt.Println("Names #2:", names2)
	*/
	// Result names2: [Agus Achmad Eko Bagus]
	// names = append(names, "rachmad", "yusuf", "edy")
	// fmt.Println("#1 Names :", names)

	// Slicing (variable name)[(start index):(previous index)]
	// names = names[1:6]
	// fmt.Println("#1 Names :", names)

	/*
		var names = make([]string, 0, 20)
		names = append(names, "Agus", "Achamd", "Eko", "Bagus")

		fmt.Println("Length:", len(names))
		fmt.Println("Capacity:", cap(names))
		fmt.Printf("Memory names: %p \n", names)

		names = append(names, "rachmad", "yusuf", "edy", "#4") // Exec: Capacity 8
		fmt.Printf("Memory names: %p \n", names)
		fmt.Println("Length:", len(names))
		fmt.Println("Capacity:", cap(names))

		names = append(names, "#5") // Exec: Capacity 16
		fmt.Printf("Memory names: %p \n", names)
		fmt.Println("Length:", len(names))
		fmt.Println("Capacity:", cap(names))
	*/

	var names = make([]string, 0)
	// names[0] = "ABC" | Invalid
	for i := range names {
		names[i] = "ABC"
	}
	fmt.Println("Names :", names)
	fmt.Println("Length:", len(names))
	fmt.Println("Capacity:", cap(names))
}
