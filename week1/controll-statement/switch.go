package main

import "fmt"

/*
	// Tipe transaksi tarik setor tabungan
	var type uint8 = 1
	if type == 1 || type == 3 || type == 4 {
		// Maka lakukan ini
	} else if type == 5 || type == 6 {
		// Maka lakukan ini #2
	}

	// Implementasi switch case
	switch type {
	default:
		// Maka lakukan ini #4
	case 1, 3, 4:
		// Maka lakukan ini #1
	case 5, 6:
		// Maka lakukan ini #2
	case 7:
		// Maka lakukan ini #3
	}
*/
func main() {
	var grade = "Grade B"

	switch grade {
	default:
		fmt.Println("Point range 0 - 69")
	case "Grade A":
		fmt.Println("Point range 90 - 100")
	case "Grade B":
		fmt.Println("Point range 80 - 89")
	case "Grade C":
		fmt.Println("Point range 70 - 79")
	}

	switch {
	case grade == "Grade A":
		fmt.Println("#2 Point range 90 - 100")
	case grade == "Grade B":
		fmt.Println("#2 Point range 80 - 89")
	}

	var point uint8 = 80
	switch {
	default:
		fmt.Println("Anda Tidak Lulus")

	case point <= 100 && point >= 70:
		fmt.Println("Selamat Anda Lulus")
	}
}
