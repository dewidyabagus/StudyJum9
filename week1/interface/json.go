package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Biodata struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

var result = `{"name": "john wick", "age": 30, "hobbies": ["berlari", "makan", "tidur"]}`

func withMap() {
	var biodata = map[string]interface{}{}

	if err := json.Unmarshal([]byte(result), &biodata); err != nil {
		log.Fatalln("unmarshal:", err.Error())
	}

	fmt.Printf("Result: %+v \n", biodata)
}

func withStruct() {
	var biodata = Biodata{}

	if err := json.Unmarshal([]byte(result), &biodata); err != nil {
		log.Fatalln("unmarshal:", err.Error())
	}

	fmt.Printf("Result: %+v \n", biodata)
}

func main() {
	withMap()

	withStruct()
}
