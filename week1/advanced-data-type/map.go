package main

import "fmt"

func main() {
	/*
		Pembuatan tipe data map bisa menggunakan
		1. Long declaration function
		   var (nama variable) map[(type data)](tipe data)
		2. Sort assignment value
			(name variable) := map[(type data)](tipe data){key: value}
		3. Menggunakan keyword make
			(name variable) := make(map[(tipe data)])(tipe data))
	*/
	/*
		// default value nil
		var biodata map[string]string
		// inisialisasi map
		biodata = map[string]string{}
		// model assignment dibawah
		biodata["name"] = "Jonh Wick"
		biodata["address"] = "Jakarta"
		biodata["hobbies"] = "football"

		fmt.Println("Bioada:", biodata)
	*/
	var biodata = map[string]string{
		"name":    "Jonh Wick",
		"address": "Jakarta",
		"hobbies": "football",
	}
	fmt.Println("Bioada:", biodata)

	biodata2 := biodata
	biodata["name"] = "New name"

	fmt.Println(biodata2)
}
