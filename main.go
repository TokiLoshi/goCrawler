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
	
	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("Done with pages now time to print them!")
	
	for normalizedURL, count := range cfg.pages {
		fmt.Printf("Page %d - item: %v", count, normalizedURL)
	}

	fmt.Println("Execution complete")
	
}