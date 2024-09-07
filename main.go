package main

import (
	"fmt"
	"os"
	"strconv"
)



func main() {
	// Takes a command line argument (URL) e.g go run . BASE_URL (root URL)

	if len(os.Args) < 4 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("incorrect max concurrency")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])

	if err != nil {
		fmt.Println("incorrect max pages")
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()
	
	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("Page %d - item: %v", count, normalizedURL)
	// }

	printReport(cfg.pages, rawBaseURL)

	fmt.Printf("Execution complete\n")
	
}