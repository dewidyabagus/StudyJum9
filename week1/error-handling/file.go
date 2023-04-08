package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.CreateTemp(".", "*")
	if err != nil {
		log.Fatal("Create temp file:", err.Error())
	}
	defer os.Remove(f.Name())
	defer f.Close()

	time.Sleep(3 * time.Second)
}
