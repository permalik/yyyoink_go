package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting yyyoink")
	data, err := os.ReadFile("data/test.json")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
