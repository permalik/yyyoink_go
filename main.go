package main

// TODO: output json from string input
// TODO: unit test
// TODO: functional test
// TODO: output json from file input
// TODO: unit test
// TODO: functional test
// TODO: pretty print output
// TODO: update name to yyy_id

import (
	"bufio"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input JSON string: ")
	// TODO: Check if Scanner more convenient: https://pkg.go.dev/bufio#Reader.ReadString
	content, err := reader.ReadString('\n')
	check(err)
	fmt.Println(content)

	// fmt.Println("Input JSON file: ")
	// data, err := os.ReadFile("data/test.json")
	// check(err)

	// os.Stdout.Write(data)
}
