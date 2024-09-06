package main

import (
	"fmt"
	"os"
)

func main() {
	// Takes a command line argument (URL) e.g go run . BASE_URL (root URL)

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	pages := make(map[string]int)
	crawlPage(rawBaseURL, rawBaseURL, pages)
	fmt.Println("Done with pages now time to print them!")
	
	for normalizedURL, count := range pages {
		fmt.Printf("Page %d - item: %v", count, normalizedURL)
	}

	fmt.Println("Execution complete")
	
}