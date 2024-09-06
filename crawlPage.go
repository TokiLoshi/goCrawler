package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// rawCurrent is the current URL 
	// rawBase is the root URL we're crawlilng 

	// First run rawCurrentURL is a copy of rawBaseURL 
	// Subsequent calls to crawlPage will change rawCurrentURL
	// rawBase will stay the same 

	// Step 1 
	// Check rawCurrentURL is on same domain as rawBaseURL
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error: crawlPage could not parse base url: %v", err)
		return 
	}
	
	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: crawlPage could not parse current url %v \n", err)
		return 
	}

	// Check the domains 
	baseDomainName := parsedBaseURL.Hostname()
	currentDomainName := parsedCurrent.Hostname()
	fmt.Printf("Base domain: %v\n", baseDomainName)
	fmt.Printf("Current domain: %v\n", currentDomainName)
	if baseDomainName != currentDomainName {
		fmt.Printf("Domains don't match")
		return
	}

	// Step 2
	// Get normalized version of rawCurrentURL 
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("issue normalizing rawCurrentURL: %v", err)
		return 
	}
	
	fmt.Println("Normalized current: ", normalizedCurrent)
	
	
	// Step 3
	// If pages map has an entry for normalized version of current
	// increment count and be done (you've crawled this page)
	// else, add entry to the pages object for the nromalized version
	if _, visited := pages[normalizedCurrent]; visited {
		pages[normalizedCurrent]++
		return 
	}

	pages[normalizedCurrent] = 1

	// Step 4 
	// Get HTML from current URL and add print statement to watch crawler 
	
	fmt.Printf("On Crawler: crawling: %s\n", rawCurrentURL)

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("issue getting html: %v", err)
		return 
	}

	// Step 5 
	// If no error from step 4 get all the URLS from the response body
	allURLs, err := getURLsFromHTML(currentHTML, rawBaseURL)
	if err != nil {
		fmt.Printf("Issue getting all the urls: %v", err)
	}
	fmt.Println("I think we got all the URLS!")
	
	for i := 0; i < len(allURLs); i ++ {
		fmt.Println(allURLs[i])
	}

	fmt.Println("Time to get recursive")
	
	// Step 6 
	// Recursively crawl each url on the page 
	// Caution testing this - add print statements 
	// if you're stuck in a loop kill it with cnt + c
	for _, webURL := range allURLs {
		fmt.Println("Crawling link: ", webURL)
		crawlPage(rawBaseURL, webURL, pages)
		fmt.Println("Done and now we have: ")
		
	}

	// Add crawlPage in main function instead of getHML
	// When complete print keys and values of pages map to console

}