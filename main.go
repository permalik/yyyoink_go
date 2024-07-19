package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Starting yyyoink")
	data, err := os.ReadFile("data/test.json")
	check(err)

	for i := range data {
		fmt.Println(string(i))
	}
}
