package main

import (
	"fmt"
	"os"
)

func main() {
	// Takes a command line argument (URL) e.g go run . BASE_URL (root URL)
	argsWithProg := os.Args[1:]

	argsLength := len(os.Args[1:])
	if argsLength < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if argsLength > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if argsLength == 1 {
		baseURL := argsWithProg[0]
		fmt.Println("starting crawl of: ", baseURL)
	}
}