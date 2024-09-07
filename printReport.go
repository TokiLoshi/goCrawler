package main

import (
	"fmt"
)

func printReport(pages map[string]int, baseURL string) {

	// Sort the pages in map and make it look good: 
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Printf("=============================\n")
	sortedPages := sortPages(pages)
	
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.links, page.url)
	}
}